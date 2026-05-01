package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pedrolucaspalma/notes-go/sequence"
)

func main() {
	p := tea.NewProgram(ApplicationModel{
		cursorFret:          0,
		cursorString:        0,
		numberOfFretsOnNeck: 12,
		tuning:              E_STANDARD_TUNING,
		neckDisplayType:     DISPLAY_EMPTY_FRETS,
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

// ===========================================================================================================================
type ApplicationModel struct {
	cursorFret          int
	cursorString        int
	numberOfFretsOnNeck int
	tuning              Tuning
	neckDisplayType     NeckDisplayType
}

func (m ApplicationModel) Init() tea.Cmd {
	return nil
}

func (m ApplicationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down":
			if m.tuning > 0 {
				m.tuning--
			} else {
				m.tuning = E_STANDARD_TUNING
			}
		case "up":
			if m.tuning < E_STANDARD_TUNING {
				m.tuning++
			} else {
				m.tuning = D_STANDARD_TUNING
			}
		case "left":
			if m.neckDisplayType > 0 {
				m.neckDisplayType--
			} else {
				m.neckDisplayType = DISPLAY_NO_ACCIDENTALS
			}
		case "right":
			if m.neckDisplayType < DISPLAY_NO_ACCIDENTALS {
				m.neckDisplayType++
			} else {
				m.neckDisplayType = DISPLAY_EMPTY_FRETS
			}
		}
	}
	return m, nil
}

func (m ApplicationModel) View() string {
	neck, err := NewGuitarNeck(m.numberOfFretsOnNeck, m.tuning, m.neckDisplayType)
	if err != nil {
		panic(err)
	}
	return neck.String()
}

// ===========================================================================================================================
type Tuning int

const (
	E_STANDARD_TUNING Tuning = iota
	DROP_D_TUNING
	D_STANDARD_TUNING
)

type NeckDisplayType int

const (
	DISPLAY_EMPTY_FRETS NeckDisplayType = iota
	DISPLAY_ALL_NOTES_AS_SHARP
	DISPLAY_ALL_NOTES_AS_FLAT
	DISPLAY_NO_ACCIDENTALS
)

type GuitarNeck struct {
	NeckStrings   []GuitarString
	Tuning        string
	NumberOfFrets int
}

func NewGuitarNeck(
	fretNums int,
	tuning Tuning,
	neckDisplayType NeckDisplayType,
) (GuitarNeck, error) {
	neckStrings := []GuitarString{}

	openNotes := getTuning(tuning)
	for _, note := range openNotes {
		guitarString, err := NewGuitarString(
			note,
			fretNums,
			neckDisplayType,
		)
		if err != nil {
			return GuitarNeck{}, fmt.Errorf("creating guitar string: %w", err)
		}
		neckStrings = append(neckStrings, guitarString)
	}

	return GuitarNeck{
		NeckStrings:   neckStrings,
		NumberOfFrets: fretNums,
	}, nil
}

func (n GuitarNeck) String() string {
	var b strings.Builder
	for _, s := range n.NeckStrings {
		// open note + nut
		fmt.Fprintf(&b, "%-2s ║", s.OpenNote)

		// The rest of the string
		b.WriteString(s.String())
	}

	// Some space at the end of the string
	b.WriteString("    ")
	for f := 1; f <= n.NumberOfFrets; f++ {
		// The number of the fret
		fmt.Fprintf(&b, "   %-2d   ", f)
	}
	return b.String()
}

func getTuning(selectedTuning Tuning) []string {
	switch selectedTuning {
	case E_STANDARD_TUNING:
		return []string{"E", "B", "G", "D", "A", "E"}
	case DROP_D_TUNING:
		return []string{"E", "B", "G", "D", "A", "D"}
	case D_STANDARD_TUNING:
		return []string{"D", "G", "C", "F", "A", "D"}
	}

	return []string{"E", "B", "G", "D", "A", "E"}
}

// ===========================================================================================================================

type GuitarString struct {
	OpenNote string
	Frets    []Fret
}

func NewGuitarString(
	openNoteStr string,
	numFrets int,
	neckDisplayType NeckDisplayType,
) (GuitarString, error) {
	seq := sequence.NewNoteSequence()

	openNote, err := seq.Find(openNoteStr)
	if err != nil {
		return GuitarString{}, fmt.Errorf("finding open note for string: %w", err)
	}

	frets := []Fret{}
	for i := range numFrets {
		if i == 0 {
			continue
		}
		currentNote := openNote.FindFromSemitones(i)

		fret := NewFret(i, currentNote, getFretDisplayType(currentNote, neckDisplayType))
		frets = append(frets, fret)
	}

	return GuitarString{
		OpenNote: openNoteStr,
		Frets:    frets,
	}, nil
}

func getFretDisplayType(
	note *sequence.Note,
	neckDisplayType NeckDisplayType,
) FretDisplayType {
	switch neckDisplayType {
	case DISPLAY_ALL_NOTES_AS_FLAT:
		return DISPLAY_AS_FLAT
	case DISPLAY_ALL_NOTES_AS_SHARP:
		return DISPLAY_AS_SHARP
	case DISPLAY_NO_ACCIDENTALS:
		if note.IsAccidental() {
			return DISPLAY_EMPTY_FRET
		}
		return DISPLAY_AS_SHARP

	default:
		return DISPLAY_EMPTY_FRET
	}
}

func (s GuitarString) String() string {
	var b strings.Builder

	for _, f := range s.Frets {
		b.WriteString(f.String())
	}

	b.WriteString("\n")

	return b.String()
}

// ===========================================================================================================================

type FretDisplayType int

const (
	DISPLAY_EMPTY_FRET FretDisplayType = iota
	DISPLAY_AS_SHARP
	DISPLAY_AS_FLAT
)

type Fret struct {
	Number      int
	Note        *sequence.Note
	DisplayType FretDisplayType
}

func NewFret(
	number int,
	note *sequence.Note,
	displayType FretDisplayType,
) Fret {
	return Fret{
		Number:      number,
		Note:        note,
		DisplayType: displayType,
	}
}

func (f Fret) String() string {
	switch f.DisplayType {
	case DISPLAY_AS_FLAT:
		return fmt.Sprintf("━━ %-2s ━┿", f.Note.NameFlat)
	case DISPLAY_AS_SHARP:
		return fmt.Sprintf("━━ %-2s ━┿", f.Note.NameSharp)
	case DISPLAY_EMPTY_FRET:
		return "━━━━━━━┿"
	default:
		return "━━━━━━━┿"
	}
}

// ===========================================================================================================================

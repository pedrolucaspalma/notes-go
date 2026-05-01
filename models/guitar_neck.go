package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/pedrolucaspalma/notes-go/constants"
)

type GuitarNeck struct {
	NeckStrings     []guitarString
	Tuning          constants.Tuning
	DisplayType     constants.NeckDisplayType
	NumberOfFrets   int
}

type TuningChangedMsg struct {
	Tuning constants.Tuning
}

type DisplayTypeChangedMsg struct {
	DisplayType constants.NeckDisplayType
}

func (n GuitarNeck) Init() tea.Cmd {
	return nil
}

func (n GuitarNeck) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TuningChangedMsg:
		n.Tuning = msg.Tuning
		n.rebuildStrings()
	case DisplayTypeChangedMsg:
		n.DisplayType = msg.DisplayType
		n.rebuildStrings()
	}
	return n, nil
}

func (n *GuitarNeck) rebuildStrings() {
	neckStrings := []guitarString{}
	openNotes := getTuning(n.Tuning)
	for _, note := range openNotes {
		guitarString, err := NewGuitarString(
			note,
			n.NumberOfFrets,
			n.DisplayType,
		)
		if err == nil {
			neckStrings = append(neckStrings, guitarString)
		}
	}
	n.NeckStrings = neckStrings
}

func (n GuitarNeck) View() tea.View {
	return tea.NewView(n.String())
}

func NewGuitarNeck(
	fretNums int,
	tuning constants.Tuning,
	neckDisplayType constants.NeckDisplayType,
) (GuitarNeck, error) {
	neckStrings := []guitarString{}

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
		Tuning:        tuning,
		DisplayType:   neckDisplayType,
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

func getTuning(selectedTuning constants.Tuning) []string {
	switch selectedTuning {
	case constants.E_STANDARD_TUNING:
		return []string{"E", "B", "G", "D", "A", "E"}
	case constants.DROP_D_TUNING:
		return []string{"E", "B", "G", "D", "A", "D"}
	case constants.D_STANDARD_TUNING:
		return []string{"D", "G", "C", "F", "A", "D"}
	}

	return []string{"E", "B", "G", "D", "A", "E"}
}

package tuicomponents

import (
	"fmt"
	"strings"

	"github.com/pedrolucaspalma/notes-go/constants"
)

type guitarNeck struct {
	NeckStrings   []guitarString
	Tuning        string
	NumberOfFrets int
}

func NewGuitarNeck(
	fretNums int,
	tuning constants.Tuning,
	neckDisplayType constants.NeckDisplayType,
) (guitarNeck, error) {
	neckStrings := []guitarString{}

	openNotes := getTuning(tuning)
	for _, note := range openNotes {
		guitarString, err := NewGuitarString(
			note,
			fretNums,
			neckDisplayType,
		)
		if err != nil {
			return guitarNeck{}, fmt.Errorf("creating guitar string: %w", err)
		}
		neckStrings = append(neckStrings, guitarString)
	}

	return guitarNeck{
		NeckStrings:   neckStrings,
		NumberOfFrets: fretNums,
	}, nil
}

func (n guitarNeck) String() string {
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

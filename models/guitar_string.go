package models

import (
	"fmt"
	"strings"

	"github.com/pedrolucaspalma/notes-go/constants"
	"github.com/pedrolucaspalma/notes-go/sequence"
)

type guitarString struct {
	OpenNote string
	Frets    []fret
}

func NewGuitarString(
	openNoteStr string,
	numFrets int,
	neckDisplayType constants.NeckDisplayType,
) (guitarString, error) {
	seq := sequence.NewNoteSequence()

	openNote, err := seq.Find(openNoteStr)
	if err != nil {
		return guitarString{}, fmt.Errorf("finding open note for string: %w", err)
	}

	frets := []fret{}
	for i := range numFrets {
		if i == 0 {
			continue
		}
		currentNote := openNote.FindFromSemitones(i)

		fret := NewFret(i, currentNote, getFretDisplayType(currentNote, neckDisplayType))
		frets = append(frets, fret)
	}

	return guitarString{
		OpenNote: openNoteStr,
		Frets:    frets,
	}, nil
}

func getFretDisplayType(
	note *sequence.Note,
	neckDisplayType constants.NeckDisplayType,
) constants.FretDisplayType {
	switch neckDisplayType {
	case constants.DISPLAY_ALL_NOTES_AS_FLAT:
		return constants.DISPLAY_AS_FLAT
	case constants.DISPLAY_ALL_NOTES_AS_SHARP:
		return constants.DISPLAY_AS_SHARP
	case constants.DISPLAY_NO_ACCIDENTALS:
		if note.IsAccidental() {
			return constants.DISPLAY_EMPTY_FRET
		}
		return constants.DISPLAY_AS_SHARP

	default:
		return constants.DISPLAY_EMPTY_FRET
	}
}

func (s guitarString) String() string {
	var b strings.Builder

	for _, f := range s.Frets {
		b.WriteString(f.String())
	}

	b.WriteString("\n")

	return b.String()
}

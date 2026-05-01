package sequence

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"strings"
)

type NoteSequence struct {
	first *Note
}

func NewNoteSequence() *NoteSequence {
	notesSharp := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	notesFlat := []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

	var firstNote *Note
	var previous *Note

	for i := range 12 {
		current := &Note{
			NameSharp: notesSharp[i],
			NameFlat:  notesFlat[i],
			Next:      nil,
			Previous:  previous,
		}

		if i == 0 {
			firstNote = current
		} else {
			previous.Next = current
			current.Previous = previous
		}
		previous = current

	}

	previous.Next = firstNote
	firstNote.Previous = firstNote

	return &NoteSequence{
		first: firstNote,
	}
}

func (s *NoteSequence) PrintAllIntervals() {
	fmt.Printf("0: Unison\n")
	fmt.Printf("1: Minor Second\n")
	fmt.Printf("2: Major Second\n")
	fmt.Printf("3: Minor Third\n")
	fmt.Printf("4: Major Third\n")
	fmt.Printf("5: Perfect Fourth\n")
	fmt.Printf("6: Augmented Fourth/Diminished Fith/Tritone\n")
	fmt.Printf("7: Perfect Fith\n")
	fmt.Printf("8: Augmented Fith/Minor Sixth\n")
	fmt.Printf("9: Major Sixth\n")
	fmt.Printf("10: Minor Seventh\n")
	fmt.Printf("11: Major Seventh\n")
	fmt.Printf("12: Perfect Octave\n")
}

func (s *NoteSequence) chromatic(start *Note) iter.Seq[*Note] {
	return func(yield func(*Note) bool) {
		curr := start
		for range 12 {
			if !yield(curr) {
				return
			}

			curr = curr.Next
		}
	}
}

func (s *NoteSequence) Find(note string) (*Note, error) {
	note = strings.ToLower(note)

	for curr := range s.chromatic(s.first) {
		if strings.ToLower(curr.NameFlat) == note || strings.ToLower(curr.NameSharp) == note {
			return curr, nil
		}
	}
	return nil, NOTE_NOT_FOUND_ERROR
}

func (s *NoteSequence) getNotesFromIntervals(root *Note, intervalsInSemitones []int) []Note {
	notes := []Note{}

	count := 0
	for curr := range s.chromatic(root) {
		if slices.Contains(intervalsInSemitones, count) {
			notes = append(notes, *curr)
		}
		count++
	}

	return notes
}

func (s *NoteSequence) findNoteAndGetNotesFromIntervals(note string, intervals []int) ([]Note, error) {
	root, err := s.Find(note)
	if err != nil {
		return nil, fmt.Errorf("finding root note: %w", err)
	}
	notes := s.getNotesFromIntervals(root, intervals)
	return notes, nil
}

var NOTE_NOT_FOUND_ERROR = errors.New("note not found")

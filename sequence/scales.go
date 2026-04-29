package sequence

// Chromatic

// Ionian
func (s *NoteSequence) MajorDiatonicScale(note string) ([]Note, error) {
	// whole, whole, half, whole, whole, whole, half
	intervalsInSemitones := []int{0, 2, 4, 5, 7, 9, 11, 12}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

func (s *NoteSequence) MinorScale(note string) ([]Note, error) {
	// whole half whole whole half whole whole
	intervalsInSemitones := []int{0, 2, 3, 5, 7, 8, 10, 12}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

func (s *NoteSequence) PentatonicMajorScale(note string) ([]Note, error) {
	majorNotes, err := s.MajorDiatonicScale(note)
	if err != nil {
		return nil, err
	}

	// 1,2,3,5,6 tonics from major scale
	return []Note{
		majorNotes[0],
		majorNotes[1],
		majorNotes[2],
		majorNotes[4],
		majorNotes[5],
	}, nil
}

// Lydian
// Mixolydian
// Melodic Minor
// Dorian
// Phrygian
// Aeolian
// Harmonic Minor
// Pentatonic Minor
// Whole Tone
// Diminished
// Locrian

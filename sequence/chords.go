package sequence

// ====================================== Triads ==============================================

// Ex C (C major)
func (s *NoteSequence) MajorChord(note string) ([]Note, error) {
	// root, major 3rd, perfect 5h
	intervalsInSemitones := []int{0, 4, 7}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

// Ex Cm (C minor)
func (s *NoteSequence) MinorChord(note string) ([]Note, error) {
	// root, minor 3rd, perfect 5h
	intervalsInSemitones := []int{0, 3, 7}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

// ====================================== Augmented/Diminished Triads Chords ==============================================
// Ex Caug C+ (C augmented)
func (s *NoteSequence) AugmentedChord(note string) ([]Note, error) {
	// root, major 3rd, augmented 5h
	intervalsInSemitones := []int{0, 4, 8}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

// Ex Cdim Cº (C diminished)
func (s *NoteSequence) DiminishedChord(note string) ([]Note, error) {
	// root, minor 3rd, diminished 5th/augmented 4th
	intervalsInSemitones := []int{0, 3, 6}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

// ====================================== 7 Chords ==============================================

// Ex C7 (C Dominanth Seventh)
// The major triad + minor seventh
func (s *NoteSequence) DominantSeventh(note string) ([]Note, error) {
	// root, minor 3rd, perfect 5h
	intervalsInSemitones := []int{0, 4, 7, 10}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

// Ex Cmaj7 (C Major Seventh)
// The major triad + major seventh
func (s *NoteSequence) MajorSeventh(note string) ([]Note, error) {
	// root, minor 3rd, perfect 5h
	intervalsInSemitones := []int{0, 4, 7, 11}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

// Ex Cmin7 (C Minor Seventh)
// The minor triad + minor seventh
func (s *NoteSequence) MinorSeventh(note string) ([]Note, error) {
	// root, minor 3rd, perfect 5h
	intervalsInSemitones := []int{0, 3, 7, 10}
	notes, err := s.findNoteAndGetNotesFromIntervals(note, intervalsInSemitones)
	return notes, err
}

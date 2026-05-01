package constants

type Tuning int

const (
	E_STANDARD_TUNING Tuning = iota
	DROP_D_TUNING
	D_STANDARD_TUNING
)

// ===========================================

type NeckDisplayType int

const (
	DISPLAY_EMPTY_FRETS NeckDisplayType = iota
	DISPLAY_ALL_NOTES_AS_SHARP
	DISPLAY_ALL_NOTES_AS_FLAT
	DISPLAY_NO_ACCIDENTALS
)

// ===========================================

type FretDisplayType int

const (
	DISPLAY_EMPTY_FRET FretDisplayType = iota
	DISPLAY_AS_SHARP
	DISPLAY_AS_FLAT
	DISPLAY_CIRCLE
)

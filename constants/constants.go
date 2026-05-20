package constants

type Tuning int

func (t Tuning) HasRotatedAllTypes() bool {
	return int(t) == len(TUNING_DISPLAY_STR_MAP)-1
}

const (
	E_STANDARD_TUNING Tuning = iota
	DROP_D_TUNING
	D_STANDARD_TUNING
)

var TUNING_DISPLAY_STR_MAP = map[Tuning]string{
	E_STANDARD_TUNING: "E Standard",
	DROP_D_TUNING:     "Drop D",
	D_STANDARD_TUNING: "D Standard",
}

// ===========================================

type NeckDisplayType int

func (n NeckDisplayType) HasRotatedAllTypes() bool {
	return int(n) == len(NECK_DISPLAY_TYPE_DISPLAY_STR_MAP)-1
}

const (
	DISPLAY_EMPTY_FRETS NeckDisplayType = iota
	DISPLAY_ALL_NOTES_AS_SHARP
	DISPLAY_ALL_NOTES_AS_FLAT
	DISPLAY_NO_ACCIDENTALS
)

var NECK_DISPLAY_TYPE_DISPLAY_STR_MAP = map[NeckDisplayType]string{
	DISPLAY_EMPTY_FRETS:        "Empty Frets",
	DISPLAY_ALL_NOTES_AS_SHARP: "Sharps",
	DISPLAY_ALL_NOTES_AS_FLAT:  "Flats",
	DISPLAY_NO_ACCIDENTALS:     "Natural Notes Only",
}

// ===========================================

type FretDisplayType int

const (
	DISPLAY_EMPTY_FRET FretDisplayType = iota
	DISPLAY_AS_SHARP
	DISPLAY_AS_FLAT
	DISPLAY_CIRCLE
)

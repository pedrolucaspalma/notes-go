package constants

type colors struct {
	PURPLE          string
	CYAN            string
	PINK            string
	DARK_BLACK_GRAY string
	SILVER_STEEL    string
	MARBLE          string
	PASTEL_RED      string
	PASTEL_YELLOW   string
	WHITE           string
}

var COLORS = colors{
	PURPLE:          "#5f5fff",
	CYAN:            "#5fffd7",
	PINK:            "#ff5faf",
	DARK_BLACK_GRAY: "#626262",
	SILVER_STEEL:    "#9BA1AB",
	MARBLE:          "#e3e0cd",
	PASTEL_RED:      "#FF746C",
	PASTEL_YELLOW:   "#FFEE8C",
	WHITE:           "#FFFFFF",
}

// ===========================================

type componentsColors struct {
	BORDER     string
	VALUE_TEXT string
	TITLE_TEXT string
	HELP_TEXT  string
	NOTE_TEXT  string

	STRING string
	NUT    string
}

var COMPONENTS_COLORS = componentsColors{
	BORDER:     COLORS.PURPLE,
	VALUE_TEXT: COLORS.CYAN,
	HELP_TEXT:  COLORS.DARK_BLACK_GRAY,
	NOTE_TEXT:  COLORS.PASTEL_RED,

	STRING: COLORS.SILVER_STEEL,
	NUT:    COLORS.MARBLE,
}

// ===========================================

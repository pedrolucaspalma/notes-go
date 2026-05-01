package constants

type colors struct {
	PURPLE          string
	CYAN            string
	PINK            string
	DARK_BLACK_GRAY string
}

var COLORS = colors{
	PURPLE:          "#5f5fff",
	CYAN:            "#5fffd7",
	PINK:            "#ff5faf",
	DARK_BLACK_GRAY: "#626262",
}

// ===========================================

type componentsColors struct {
	BORDER     string
	VALUE_TEXT string
	TITLE_TEXT string
	HELP_TEXT  string
}

var COMPONENTS_COLORS = componentsColors{
	BORDER:     COLORS.PURPLE,
	VALUE_TEXT: COLORS.CYAN,
	HELP_TEXT:  COLORS.DARK_BLACK_GRAY,
}

// ===========================================

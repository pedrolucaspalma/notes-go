package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/pedrolucaspalma/notes-go/constants"
)

type Menu struct {
	Tuning      constants.Tuning
	DisplayType constants.NeckDisplayType
}

func NewMenuModel() Menu {
	return Menu{
		Tuning:      constants.E_STANDARD_TUNING,
		DisplayType: constants.DISPLAY_EMPTY_FRETS,
	}
}

func (m Menu) Init() tea.Cmd {
	return nil
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TuningChangedMsg:
		m.Tuning = msg.Tuning
	case DisplayTypeChangedMsg:
		m.DisplayType = msg.DisplayType
	}
	return m, nil
}

func (m Menu) View() tea.View {
	var b strings.Builder

	tuningStr := ""
	switch m.Tuning {
	case constants.E_STANDARD_TUNING:
		tuningStr = "E Standard"
	case constants.DROP_D_TUNING:
		tuningStr = "Drop D"
	case constants.D_STANDARD_TUNING:
		tuningStr = "D Standard"
	}

	displayStr := ""
	switch m.DisplayType {
	case constants.DISPLAY_EMPTY_FRETS:
		displayStr = "Empty Frets"
	case constants.DISPLAY_ALL_NOTES_AS_SHARP:
		displayStr = "Sharps"
	case constants.DISPLAY_ALL_NOTES_AS_FLAT:
		displayStr = "Flats"
	case constants.DISPLAY_NO_ACCIDENTALS:
		displayStr = "Natural Notes Only"
	}

	b.WriteString(fmt.Sprintf("Current Tuning: %s (↑/↓ to change)\n", tuningStr))
	b.WriteString(fmt.Sprintf("Fret Display: %s (←/→ to change)\n", displayStr))
	b.WriteString(strings.Repeat("-", 40) + "\n\n")

	return tea.NewView(b.String())
}

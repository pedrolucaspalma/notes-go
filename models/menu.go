package models

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(constants.COLORS.PINK))

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(constants.COMPONENTS_COLORS.VALUE_TEXT))

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(constants.COLORS.DARK_BLACK_GRAY)).
		Italic(true)

	tuningLine := lipgloss.JoinHorizontal(lipgloss.Left,
		titleStyle.Render("Current Tuning: "),
		valueStyle.Render(tuningStr),
		helpStyle.Render(" (↑/↓ to change)"),
	)

	displayLine := lipgloss.JoinHorizontal(lipgloss.Left,
		titleStyle.Render("Fret Display:   "),
		valueStyle.Render(displayStr),
		helpStyle.Render(" (←/→ to change)"),
	)

	b.WriteString(tuningLine + "\n")
	b.WriteString(displayLine + "\n")
	b.WriteString(strings.Repeat("-", 40) + "\n\n")

	return tea.NewView(b.String())
}

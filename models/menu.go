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

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(constants.COLORS.PINK))

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(constants.COMPONENTS_COLORS.VALUE_TEXT))

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(constants.COLORS.DARK_BLACK_GRAY)).
		Italic(true)

	tuningStr := constants.TUNING_DISPLAY_STR_MAP[m.Tuning]
	tuningLine := lipgloss.JoinHorizontal(lipgloss.Left,
		titleStyle.Render("Current Tuning: "),
		valueStyle.Render(tuningStr),
		helpStyle.Render(" (↑/↓ to change)"),
	)

	displayStr := constants.NECK_DISPLAY_TYPE_DISPLAY_STR_MAP[m.DisplayType]
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

package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/pedrolucaspalma/notes-go/constants"
	"github.com/pedrolucaspalma/notes-go/models"
)

func main() {
	guitarNeck, err := models.NewGuitarNeck(
		12,
		constants.E_STANDARD_TUNING,
		constants.DISPLAY_EMPTY_FRETS,
	)
	if err != nil {
		panic(err)
	}

	applicationMenu := models.NewMenuModel()

	p := tea.NewProgram(ApplicationModel{
		models: applicationSubmodels{
			guitarNeck: guitarNeck,
			menu:       applicationMenu,
		},

		cursorFret:          0,
		cursorString:        0,
		numberOfFretsOnNeck: 12,
		tuning:              constants.E_STANDARD_TUNING,
		neckDisplayType:     constants.DISPLAY_EMPTY_FRETS,
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

type applicationSubmodels struct {
	guitarNeck models.GuitarNeck
	menu       models.Menu
}

type ApplicationModel struct {
	models applicationSubmodels

	width  int
	height int

	cursorFret          int
	cursorString        int
	cursorMenu          int
	numberOfFretsOnNeck int
	tuning              constants.Tuning
	neckDisplayType     constants.NeckDisplayType
}

func (m ApplicationModel) Init() tea.Cmd {
	return nil
}

func (m ApplicationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	updatedNeck, cmd := m.models.guitarNeck.Update(msg)
	m.models.guitarNeck = updatedNeck.(models.GuitarNeck)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	updatedMenu, cmd := m.models.menu.Update(msg)
	m.models.menu = updatedMenu.(models.Menu)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "t":
			if int(m.tuning) == len(constants.TUNING_DISPLAY_STR_MAP) {
				m.tuning = 0
			} else {
				m.tuning++
			}
			cmds = append(cmds, func() tea.Msg { return models.TuningChangedMsg{Tuning: m.tuning} })
		case "f":
			if int(m.neckDisplayType) == len(constants.NECK_DISPLAY_TYPE_DISPLAY_STR_MAP) {
				m.neckDisplayType = 0
			} else {
				m.neckDisplayType++
			}
			cmds = append(cmds, func() tea.Msg { return models.DisplayTypeChangedMsg{DisplayType: m.neckDisplayType} })
		}
	}
	return m, tea.Batch(cmds...)
}

func (m ApplicationModel) View() tea.View {
	menuView := m.models.menu.View()
	neckView := m.models.guitarNeck.View()

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		menuView.Content,
		"", // spacing
		neckView.Content,
	)

	appStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(constants.COMPONENTS_COLORS.BORDER)).
		Padding(1, 4)

	rendered := appStyle.Render(content)

	// Center horizontally and vertically
	centered := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		rendered,
	)

	finalView := tea.NewView(centered)
	finalView.AltScreen = true

	return finalView
}

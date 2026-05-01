package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
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
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down":
			if m.tuning > 0 {
				m.tuning--
			} else {
				m.tuning = constants.D_STANDARD_TUNING
			}
			cmds = append(cmds, func() tea.Msg { return models.TuningChangedMsg{Tuning: m.tuning} })
		case "up":
			if m.tuning < constants.D_STANDARD_TUNING {
				m.tuning++
			} else {
				m.tuning = constants.E_STANDARD_TUNING
			}
			cmds = append(cmds, func() tea.Msg { return models.TuningChangedMsg{Tuning: m.tuning} })
		case "left":
			if m.neckDisplayType > 0 {
				m.neckDisplayType--
			} else {
				m.neckDisplayType = constants.DISPLAY_NO_ACCIDENTALS
			}
			cmds = append(cmds, func() tea.Msg { return models.DisplayTypeChangedMsg{DisplayType: m.neckDisplayType} })
		case "right":
			if m.neckDisplayType < constants.DISPLAY_NO_ACCIDENTALS {
				m.neckDisplayType++
			} else {
				m.neckDisplayType = constants.DISPLAY_EMPTY_FRETS
			}
			cmds = append(cmds, func() tea.Msg { return models.DisplayTypeChangedMsg{DisplayType: m.neckDisplayType} })
		}
	}
	return m, tea.Batch(cmds...)
}

func (m ApplicationModel) View() tea.View {
	var b strings.Builder

	menuView := m.models.menu.View()
	neckView := m.models.guitarNeck.View()
	b.WriteString(menuView.Content)
	b.WriteString(neckView.Content)

	finalView := tea.NewView(b.String())

	return finalView
}

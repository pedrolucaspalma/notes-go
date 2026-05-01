package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/pedrolucaspalma/notes-go/constants"
	"github.com/pedrolucaspalma/notes-go/tuicomponents"
)

func main() {
	p := tea.NewProgram(ApplicationModel{
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

type ApplicationModel struct {
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down":
			if m.tuning > 0 {
				m.tuning--
			} else {
				m.tuning = constants.E_STANDARD_TUNING
			}
		case "up":
			if m.tuning < constants.E_STANDARD_TUNING {
				m.tuning++
			} else {
				m.tuning = constants.D_STANDARD_TUNING
			}
		case "left":
			if m.neckDisplayType > 0 {
				m.neckDisplayType--
			} else {
				m.neckDisplayType = constants.DISPLAY_NO_ACCIDENTALS
			}

		case "right":
			if m.neckDisplayType < constants.DISPLAY_NO_ACCIDENTALS {
				m.neckDisplayType++
			} else {
				m.neckDisplayType = constants.DISPLAY_EMPTY_FRETS
			}
		}
	}
	return m, nil
}

func (m ApplicationModel) View() tea.View {
	var b strings.Builder

	neck, err := tuicomponents.NewGuitarNeck(m.numberOfFretsOnNeck, m.tuning, m.neckDisplayType)
	if err != nil {
		panic(err)
	}

	b.WriteString(neck.String())
	return tea.NewView(b.String())
}

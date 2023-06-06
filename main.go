package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	m := initialModel()

	rows := []*stickers.FlexBoxRow{
		m.flexBox.NewRow().AddCells([]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(1, 1).SetContent(m.currentTime),
			stickers.NewFlexBoxCell(1, 1).SetContent(lipgloss.NewStyle().Align(lipgloss.Right).Render("Battery: 100%")),
		}),
		m.flexBox.NewRow().AddCells([]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(1, 1).SetContent(m.currentDate),
		}),
		m.flexBox.NewRow().AddCells([]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(1, 11),
		}),
	}

	m.flexBox.AddRows(rows)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type Model struct {
	currentTime string
	currentDate string
	flexBox     *stickers.FlexBox
}

func initialModel() Model {
	return Model{
		currentTime: "12:00 PM",
		currentDate: "MON, JAN 1",
		flexBox:     stickers.NewFlexBox(0, 0),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	t := time.Now()
	m.currentTime = t.Format("3:04 PM")
	m.currentDate = strings.ToUpper(t.Format("Mon, Jan _2"))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.flexBox.SetWidth(msg.Width)
		m.flexBox.SetHeight(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	// topLeft := lipgloss.NewStyle().Align(lipgloss.Left).Width(40 / 2).Render(m.currentTime + "\n" + m.currentDate)
	// topRight := lipgloss.NewStyle().Align(lipgloss.Right).Width(40 / 2).Render("Battery: 100%")
	// menuBar := lipgloss.JoinHorizontal(lipgloss.Top, topLeft, topRight)

	// // Menu with 4 items
	// menu := lipgloss.NewStyle().Align(lipgloss.Center).Width(40).Render(
	// 	" Home   Documents   Downloads   Pictures   Movies   Music   Applications   Terminal   Log Out",
	// )

	return m.flexBox.Render()
}

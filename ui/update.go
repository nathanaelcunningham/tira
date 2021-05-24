package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tira/config"
	"github.com/nathanaelcunningham/tira/pane"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cfg := config.GetConfig()
		if !m.Ready {
			m.ServersPane = pane.NewModel(msg.Width, msg.Height)
			m.ServersPane.SetContent(cfg.User.Username)
			m.Ready = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

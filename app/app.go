package app

import (
	"log"
    "os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tira/ui"
) 
func Run() {

    m := ui.NewModel()
    p := tea.NewProgram(m)
    p.EnterAltScreen()

    if err := p.Start(); err != nil {
        log.Fatal("Failed to start tira", err)
        os.Exit(1)
    }

}

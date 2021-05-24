package ui

// import (
// 	"github.com/charmbracelet/lipgloss"
// )
func (m Model) View() string {

    topPane := m.ServersPane.View()

    return topPane 
}

package pane

import (
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Content string
	Width   int
	Height  int
}

func NewModel(width, height int) Model {
	m := Model{
		Content: "",
		Width:   width,
		Height:  height,
	}

	return m
}

func (m *Model) SetContent(content string) {
	m.Content = content
}

func (m Model) View() string {
    border := lipgloss.NormalBorder()
    width := m.Width - lipgloss.Width(border.Left + border.Right)

	return lipgloss.NewStyle().
            Border(border).
            Width(width).
            Height(2).
            Render(m.Content)
}

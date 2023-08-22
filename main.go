package main

import (
	"github.com/CalebJohnHunt/stacker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type base struct {
	tea.Model
}

var wrapper = lipgloss.NewStyle()

func (b *base) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		wrapper = wrapper.Width(msg.Width).MaxHeight(msg.Height)
		return b, nil
	}
	var c tea.Cmd
	b.Model, c = b.Model.Update(msg)
	return b, c
}

func (b *base) View() string {
	return wrapper.Render(b.Model.View())
}

func main() {
	st := stacker.NewStacker(&HottestPosts{})
	_, _ = tea.NewProgram(&base{st}, tea.WithAltScreen()).Run()
}

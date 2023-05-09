package main

import (
	"fmt"
	"lanc/models"
	"os"
	"strings"

	"github.com/CalebJohnHunt/stacker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SinglePost struct {
	shortId     string
	post        models.Post
	doneLoading bool
}

func (s *SinglePost) Init() tea.Cmd {
	return func() tea.Msg {
		m, err := getPost(s.shortId)
		if err != nil {
			return err
		}
		return m
	}
}

func (s *SinglePost) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return s, stacker.PopScene()
		case "w":
			f, _ := os.OpenFile("log", os.O_APPEND|os.O_CREATE, os.FileMode(0o777))
			defer f.Close()
			f.Write([]byte(fmt.Sprintf("%v", s.post)))
		}
	case models.Post:
		s.post = msg
		s.doneLoading = true
	case error:
		fmt.Println(msg)
		s.doneLoading = true
	}
	return s, nil
}

var (
	descriptionStyle                  = lipgloss.NewStyle().Border(lipgloss.DoubleBorder())
	titleStyle                        = lipgloss.NewStyle().Border(lipgloss.DoubleBorder())
	commentStyles    []lipgloss.Style = []lipgloss.Style{
		lipgloss.NewStyle().Foreground(lipgloss.Color("#f00")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#0f0")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#00f")),
	}
)

func (s *SinglePost) View() string {
	if !s.doneLoading {
		return "Loading..."
	}
	sb := strings.Builder{}
	sb.WriteString(descriptionStyle.Render(s.post.Title))
	sb.WriteByte('\n')
	if s.post.DescriptionPlain != "" {
		sb.WriteString(descriptionStyle.Render(s.post.DescriptionPlain))
		sb.WriteByte('\n')
	}
	sb.WriteString(titleStyle.Render(s.post.URL))
	sb.WriteByte('\n')

	f, _ := os.OpenFile("log", os.O_APPEND|os.O_CREATE, os.FileMode(0o777))
	defer f.Close()

	for i, comment := range s.post.Comments {
		text := strings.ReplaceAll(comment.CommentPlain, "\r", "")
		// text := strings.ReplaceAll(comment.CommentPlain, "\n", "")
		sb.WriteString(commentStyles[i%len(commentStyles)].
			Width(100).
			Border(lipgloss.NormalBorder()).
			PaddingLeft(comment.IndentLevel * 2).
			Render(text[:func() int {
				if 50 < len(text) {
					return 50
				}
				return len(text)
			}()] + "..."))
		// Render(comment.CommentPlain))
		sb.WriteByte('\n')
		f.WriteString(text)
		f.Write([]byte{'\n'})
	}

	return sb.String()
}

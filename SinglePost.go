package main

import (
	"fmt"
	"lanc/dto"
	"math"
	"strings"

	"github.com/CalebJohnHunt/stacker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SinglePost struct {
	shortId           string
	post              dto.Post
	doneLoading       bool
	selectedComment   int
	collapsedComments map[int]bool
	// Comments hidden due to their ancestor being collapsed.
	hiddenComments map[int]bool
}

func (s *SinglePost) Init() tea.Cmd {
	if s.collapsedComments == nil {
		s.collapsedComments = map[int]bool{}
	}
	if s.hiddenComments == nil {
		s.hiddenComments = map[int]bool{}
	}
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
		case "j":
			for i := s.selectedComment + 1; i < len(s.post.Comments); i++ {
				if !s.hiddenComments[i] {
					s.selectedComment = i
					break
				}
			}
		case "k":
			for i := s.selectedComment - 1; i >= 0; i-- {
				if !s.hiddenComments[i] {
					s.selectedComment = i
					break
				}
			}
		case "enter", "l", "h":
			s.collapsedComments[s.selectedComment] = !s.collapsedComments[s.selectedComment]
			for i := s.selectedComment + 1; i < len(s.post.Comments); i++ {
				if s.post.Comments[i].IndentLevel <= s.post.Comments[s.selectedComment].IndentLevel {
					break
				}
				s.hiddenComments[i] = s.collapsedComments[s.selectedComment]
				// I'm a comment. If I'm collapsed but I've just become unhidden, I don't want my children to become unhidden.
				if !s.hiddenComments[i] && s.collapsedComments[i] {
					j := i + 1
					for ; j < len(s.post.Comments) && s.post.Comments[j].IndentLevel > s.post.Comments[i].IndentLevel; j++ {
						i = j
					}
				}
			}
		}
	case dto.Post:
		s.post = msg
		s.doneLoading = true
	case error:
		fmt.Println(msg)
		s.doneLoading = true
	}
	return s, nil
}

var (
	descriptionStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder())
	titleStyle       = lipgloss.NewStyle().Border(lipgloss.DoubleBorder())
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
	sb.WriteString(titleStyle.Render(s.post.CommentsURL))
	sb.WriteByte('\n')
	sb.WriteString(titleStyle.Render(s.post.URL))
	sb.WriteByte('\n')

	s.renderComments(&sb)

	return wrapper.Render(sb.String())
}

func (s *SinglePost) renderComments(sb *strings.Builder) {
	skipIndentedMoreThanThis := math.MaxInt
	for i, comment := range s.post.Comments {
		if comment.IndentLevel > skipIndentedMoreThanThis {
			continue
		}
		skipIndentedMoreThanThis = math.MaxInt
		if s.collapsedComments[i] {
			skipIndentedMoreThanThis = comment.IndentLevel
		}
		text := strings.ReplaceAll(comment.CommentPlain, "\r", "")
		collapsed := "-"
		if s.collapsedComments[i] {
			collapsed = "+"
		}
		if i == s.selectedComment {
			collapsed = lipgloss.NewStyle().Background(lipgloss.Color("#ffffff")).Foreground(lipgloss.Color("#000000")).Render(collapsed)
		}

		sb.WriteString(fmt.Sprintf("%s[%s] [%d] %s (%s) (%s)\n",
			strings.Repeat("  ", comment.IndentLevel-1),
			collapsed,
			comment.Score,
			comment.CommentingUser.Username,
			comment.CreatedAt.Format("2006-01-02 03:04PM"),
			comment.ShortIDURL))
		if !s.collapsedComments[i] {
			sb.WriteString(lipgloss.JoinHorizontal(
				lipgloss.Top,
				"  ",
				strings.Repeat("  ", comment.IndentLevel),
				wrapper.Copy().Width(wrapper.GetWidth()-(comment.IndentLevel-1)*2-4).Render(text)))

			sb.WriteByte('\n')
		}
	}
}

package main

import (
	"fmt"
	"lanc/dto"
	"strings"

	"github.com/CalebJohnHunt/stacker"
	tea "github.com/charmbracelet/bubbletea"
)

type HottestPosts struct {
	posts            []dto.ShortPost
	curPage          int
	hoveredPostIndex int
}

func (h *HottestPosts) Init() tea.Cmd {
	if h.curPage < 1 {
		h.curPage = 1
	}
	return h.loadPosts()
}

func (h *HottestPosts) loadPosts() tea.Cmd {
	return func() tea.Msg {
		posts, err := getHottest(h.curPage)
		if err != nil {
			return err
		}
		a := make([]dto.ShortPost, 0, len(h.posts)+len(posts))
		a = append(a, h.posts...)
		a = append(a, posts...)
		return a
	}
}

func (h *HottestPosts) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return h, tea.Quit
		case "down", "j":
			h.hoveredPostIndex = min(len(h.posts)-1, h.hoveredPostIndex+1)
		case "up", "k":
			h.hoveredPostIndex = max(0, h.hoveredPostIndex-1)
		case "l":
			h.curPage++
			return h, h.loadPosts()
		case "enter":
			return h, stacker.AddScene(&SinglePost{shortId: h.posts[h.hoveredPostIndex].ShortID})
		}
	case error:
		fmt.Println(msg)
	case []dto.ShortPost:
		h.posts = msg
	}
	return h, nil
}

func (h *HottestPosts) View() string {
	if h.posts == nil {
		return "Waiting..."
	}
	var sb strings.Builder
	for i, post := range h.posts {
		if h.hoveredPostIndex == i {
			sb.WriteString("👉")
		} else {
			sb.WriteString("  ")
		}
		sb.WriteString(fmt.Sprintf("%3d ", post.Score))
		sb.WriteString(post.Title)
		sb.WriteString(fmt.Sprintf(" (by: %s)", post.Username))
		sb.WriteByte('\n')
	}
	return sb.String()
}

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
	return func() tea.Msg {
		posts, err := getHottest(h.curPage)
		if err != nil {
			return err
		}
		return posts
	}
}

func (h *HottestPosts) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return h, tea.Quit
		case "down":
			h.hoveredPostIndex++
		case "up":
			h.hoveredPostIndex--
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
		sb.WriteString(post.Title)
		sb.WriteByte('\n')
	}
	return sb.String()
}

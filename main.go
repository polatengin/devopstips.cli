package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"net/http"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glow/utils"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/urfave/cli"
)

type BlogPost struct {
	Title       string
	Description string
	Url         string
	Path        string
	Date        string
}

var selected *BlogPost

func getStringFromUri(url string) string {
	resp, _ := http.Get(url)

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("HTTP status %d", resp.StatusCode))
	}

	in := resp.Body

	defer in.Close()

	b, _ := ioutil.ReadAll(in)

	lines := strings.Split(string(utils.RemoveFrontmatter(b)), "\n")

	content := ""

	for i, s := range lines {
		content += strings.TrimSpace(s)
		if i+1 < len(lines) {
			content += "\n"
		}
	}

	return content
}

func getListOfBlogPosts() []BlogPost {
	resp, err := http.Get("https://devopstips.net/api/posts.json")
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	var blogPostList []BlogPost
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &blogPostList); err != nil {
		return nil
	}

	return blogPostList
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			tea.Quit()

			e, ok := m.selected[m.cursor]

			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

			log.Printf("\n\n---\n\n%v\n\n---\n\n", e)
		case "esc":
		}
	}

	return m, nil
}
func (m model) View() string {
	s := ""

	selected = nil

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		if _, ok := m.selected[i]; ok {
			selected = &choice

			displayBlogPost(*selected)
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	return s
}
func main() {
	fmt.Print("\033[H\033[2J")

	blogPostList := getListOfBlogPosts()
}

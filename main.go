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


func getStringFromUri(url string) string {
	resp, _ := http.Get(url)

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("HTTP status %d", resp.StatusCode))
	}

	in := resp.Body

	defer in.Close()
	content := ""

	for i, s := range lines {
		content += strings.TrimSpace(s)
		if i+1 < len(lines) {
			content += "\n"
		}
	}

	return content
}

func main() {
}

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

func main() {
}

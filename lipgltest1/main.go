package main

// Import stetements

import (
	"strings"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Style Definitions

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url        = lipgloss.NewStyle().Foreground(special).Render
	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss")

	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)
)

// Main function

func main() {
	doc := strings.Builder{}
	var (
		title strings.Builder
	)

	desc := lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("Style Definitions for Nice Terminal Layouts"),
		infoStyle.Render("From Charm"+divider+url("https://github.com/charmbracelet/lipgloss")),
	)

	row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
	doc.WriteString(row + "\n\n")

	
}

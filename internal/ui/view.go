package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)


func (m SystemModel) View() string {
	if m.Quitting {
		return "Quitting"
	}
	
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.ViewHeader(),
		// m.ViewStats(),
		// m.ViewFooter(),
	)
}

func (m SystemModel) ViewHeader() string {
	castFrames := []string{"*", "+", "x", "o"}
	projectileFrames := []string{"·", "•", "◦", "●", "◉", "✶"}

	castFrame := castFrames[int(time.Now().UnixMilli()/180)%len(castFrames)]
	projectileFrame := projectileFrames[int(time.Now().UnixMilli()/90)%len(projectileFrames)]

	wizardASCII := []string{
		fmt.Sprintf("              %s", castFrame),
		"            .-.",
		"           (o o)",
		"           | O \\",
		fmt.Sprintf("            \\   \\_____%s===>", projectileFrame),
		"             `~~~~~`",
		"             /|_|_|\\",
		"            /_/   \\_\\",
	}

	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")).Render("Wizard CLI")
	tagline := lipgloss.NewStyle().Foreground(lipgloss.Color("110")).Render("Wizard shows you everything automagically")
	spell := lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Render("Casting system scan...")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		tagline,
		"",
		lipgloss.NewStyle().Foreground(lipgloss.Color("141")).Render(strings.Join(wizardASCII, "\n")),
		"",
		spell,
	)
}

// func (m SystemModel) ViewStats() string { }

// func (m SystemModel) ViewFooter() string { }

// func (m SystemModel) Leave() { }

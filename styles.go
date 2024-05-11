package main

import "github.com/charmbracelet/lipgloss"

func inactiveBorder() lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = border.MiddleBottom
	border.BottomRight = border.MiddleBottom
	return border
}

func activeBorder() lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.Bottom = " "
	tmp := border.BottomLeft
	border.BottomLeft = border.BottomRight
	border.BottomRight = tmp
	return border
}

var (
	border = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	header = lipgloss.NewStyle().Bold(true)

  bold = lipgloss.NewStyle().Bold(true)
  italic = lipgloss.NewStyle().Italic(true)

	inactiveTab = lipgloss.NewStyle().BorderStyle(inactiveBorder())
	activeTab   = lipgloss.NewStyle().BorderStyle(activeBorder())
)

var projectList = lipgloss.NewStyle().MarginTop(1).MarginLeft(2)

var (
	zigStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	goStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
	rustStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	cStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	cppStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	dartStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	swiftStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	pythonStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
)

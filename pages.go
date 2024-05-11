package main

import tea "github.com/charmbracelet/bubbletea"

type Page struct {
	name    string
	content tea.Model
}

var (
	pages = []Page{
		{"About", AboutMe{}},
		{"Projects", Projects{}},
	}
	currentPage = 0
)

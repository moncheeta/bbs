package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "github.com/charmbracelet/bubbles/key"
  "github.com/charmbracelet/bubbles/help"
)

var width = 80
var height = 25

var tabBar = ""
var content = ""
var helpBar = ""

type keyMap struct{
  Quit key.Binding
  Previous key.Binding
  Next key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
  return []key.Binding{k.Quit, k.Previous, k.Next}
}

func (k keyMap) FullHelp() [][]key.Binding {
  return [][]key.Binding{k.ShortHelp()}
}

var keys = keyMap{
  Quit: key.NewBinding(
    key.WithKeys("q", "ctrl+c"),
    key.WithHelp("q", "quit"),
  ),
  Previous: key.NewBinding(
    key.WithKeys("shift+tab", "left"),
    key.WithHelp("shift+tab/←", "previous tab"),
  ),
  Next: key.NewBinding(
    key.WithKeys("tab", "right"),
    key.WithHelp("tab/→", "next tab"),
  ),
}

type BBS struct{
  help help.Model
}

func newBBS() BBS {
  return BBS{
    help: help.New(),
  }
}

func (m BBS) Init() tea.Cmd {
	return pages[currentPage].content.Init()
}

func (m BBS) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width = msg.Width
    height = msg.Height
	case tea.KeyMsg:
    switch {
    case key.Matches(msg, keys.Quit):
			return m, tea.Quit
    case key.Matches(msg, keys.Previous):
			if currentPage-1 >= 0 {
				currentPage -= 1
			} else {
				currentPage = len(pages) - 1
			}
      return m, pages[currentPage].content.Init()
		case key.Matches(msg, keys.Next):
			if currentPage+1 <= len(pages)-1 {
				currentPage += 1
			} else {
				currentPage = 0
			}
      return m, pages[currentPage].content.Init()
		}
	}
	_, cmd := pages[currentPage].content.Update(msg)
	return m, cmd
}

func (m BBS) View() string {
	tabs := make([]string, len(pages))
	for i, page := range pages {
		tab := page.name
		if i != currentPage {
			tab = inactiveTab.Render(tab)
		} else {
			tab = activeTab.Render(tab)
		}
		tabs[i] = tab
	}
	tabBar = lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	tabBarWidth := lipgloss.Width(tabBar)
	if tabBarWidth < width {
		height := lipgloss.Height(tabBar)
		lining := lipgloss.PlaceVertical(height, lipgloss.Bottom, strings.Repeat(lipgloss.RoundedBorder().Bottom, width-tabBarWidth))
		tabBar = lipgloss.JoinHorizontal(lipgloss.Top, tabBar, lining)
	}

  content = pages[currentPage].content.View()

  helpBar = lipgloss.PlaceVertical(height - lipgloss.Height(tabBar) - lipgloss.Height(content), lipgloss.Bottom, m.help.View(keys))

	return lipgloss.JoinVertical(lipgloss.Top,
		tabBar,
    content,
    helpBar,
	)
}

func runUI() {
	_, _ = tea.NewProgram(newBBS(), tea.WithAltScreen()).Run()
}

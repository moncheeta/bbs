package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type Project struct{
  title string
  desc string
  src string
}

func (p Project) View() string {
  return lipgloss.JoinVertical(lipgloss.Top,
    header.Render(p.title),
    wordwrap.String(p.desc, width - 4),
    "",
    "source code: " + p.src,
  )
}

var projects = []Project{
  {
    "BBS",
    "This is my Bulletin Board System(BBS). Here, you can learn about me and my projects.",
    "not available yet",
  },
  {
    "CSC Website",
    "I made a website for Computer Science Club(CSC) at Prospect High School during a mini-hackathon. The website displays the date and time of the next meeting, recent annocements, members, and any projects those members want to share.",
    "https://github.com/moncheeta/computer_science_club.git",
  },
  // {
  //   "Chess Rewinder",
  //   "Played a chess game but want to wind it back? Use Chess Rewinder. You are able to record and replay your games. With your new insight, you will be able to dominate your opponents!",
  //   "not available yet",
  // },
  // {
  //   "Fitness Tracker",
  //   "Want an easy way to keep track of your progress and your workouts? Try Fitness Tracker. Keep track of your weight, personal records, and any workouts you've devised.",
  //   "not available yet",
  // },
  {
    "Veteran Donation Website",
    "For my first ever hackathon, I decided to help out veterans through donations. It interacts with Venmo's APIs for users to donate.",
    "https://github.com/moncheeta/veteran_donations.git",
  },
}

type Projects struct{}

func (m Projects) Init() tea.Cmd {
  return nil
}

func (m Projects) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, nil
}

func (m Projects) View() string {
  content := strings.Builder{}
  for _, project := range projects {
    _, _ = content.WriteString(project.View() + "\n\n")
  }
  return projectList.Render(content.String())
}

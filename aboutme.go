package main

import (
	"os"
	"strings"

	aic "github.com/TheZoraiz/ascii-image-converter/aic_package"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutMe struct{}

func (m AboutMe) Init() tea.Cmd {
	return nil
}

func (m AboutMe) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m AboutMe) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		border.Render(lipgloss.JoinVertical(lipgloss.Top,
			m.Picture(),
			m.ContactInfo(),
		)),
		" ",
		m.About(),
	)
}

func (m AboutMe) Picture() string {
	flags := aic.DefaultFlags()
	flags.Dimensions = []int{50, 25}
	flags.Colored = true
	picture, err := aic.Convert("./assets/monkey.jpg", flags)
	if err != nil {
		return "no picture"
	}
	return picture
}

func (m AboutMe) ContactInfo() string {
	return `alias: moncheeta
email: moncheeta@prime8.dev
github: https://github.com/moncheeta
linkedin: https://www.linkedin.com/in/damian-myrda`
}

func (m AboutMe) About() string {
	content := strings.Builder{}

	name, _ := os.ReadFile("./assets/name.txt")
	_, _ = content.WriteString(string(name))
	_, _ = content.WriteString(`* ` + bold.Render("monkeys") + ` are his favorite animal
* lives in the ` + bold.Render("terminal") + `
* loves ` + bold.Render("systems") + ` and ` + bold.Render("backend") + ` programming

programs in:
`)
	_, _ = content.WriteString("* " + zigStyle.Render("zig") + "\n")
	_, _ = content.WriteString("* " + goStyle.Render("go") + "\n")
	_, _ = content.WriteString("* " + rustStyle.Render("rust") + "\n")
	_, _ = content.WriteString("* " + cStyle.Render("c") + "\n")
	_, _ = content.WriteString("* " + cppStyle.Render("c++") + "\n")
	_, _ = content.WriteString("* " + dartStyle.Render("dart(flutter)") + "\n")
	_, _ = content.WriteString("* " + swiftStyle.Render("swift(swiftUI)") + "\n")
	_, _ = content.WriteString("* and of course, " + pythonStyle.Render("python") + "\n\n")

	_, _ = content.WriteString(WakatimeStats())

	return content.String()
}

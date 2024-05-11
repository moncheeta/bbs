package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	resty "github.com/go-resty/resty/v2"
)

var client = resty.New()

type WakatimeResponse struct {
	Data WakatimeData `json:"data"`
}

type WakatimeData struct {
	Languages []LanguageStats `json:"languages"`
	TotalTime string          `json:"human_readable_total_including_other_language"`
}

type LanguageStats struct {
	Name        string  `json:"name"`
	TotalTime   string  `json:"text"`
	PercentTime float64 `json:"percent"`
}

var stats = getStats()

func getStats() WakatimeData {
	r, err := client.R().
		SetHeader("Accept", "application/json").
		Get("https://wakatime.com/api/v1/users/moncheeta/stats/last_7_days")
	if err != nil {
		return WakatimeData{}
	}
	wr := WakatimeResponse{}
	if err = json.Unmarshal(r.Body(), &wr); err != nil {
		return WakatimeData{}
	}
	return wr.Data
}

func WakatimeStats() string {
	languages := stats.Languages
	total := stats.TotalTime

	names := strings.Builder{}
	time := strings.Builder{}
	for i, language := range languages {
		if i >= 5 {
			break
		}
		_, _ = names.WriteString("* " + language.Name + "\n")
		_, _ = time.WriteString(
			language.TotalTime +
				"(" + strconv.FormatFloat(language.PercentTime, 'f', 1, 64) + "%" + ")" +
				"\n")
	}
	return lipgloss.JoinVertical(lipgloss.Top,
		"activity:",
		lipgloss.JoinHorizontal(lipgloss.Top, names.String(), " ", time.String()),
		"total: "+total,
	)
}

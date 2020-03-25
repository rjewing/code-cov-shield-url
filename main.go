package main

import (
	"flag"
	"fmt"
	"net/url"
)

const (
	red = iota
	orange
	yellow
	yellowgreen
	green
	brightgreen
	baseURL string = "https://img.shields.io/badge/"
)

var (
	covPtr *int           = flag.Int("coverage", -1, "code coverage as a percentage")
	colors map[int]string = map[int]string{
		-1:          "unknown",
		red:         "red",
		orange:      "orange",
		yellow:      "yellow",
		yellowgreen: "yellowgreen",
		green:       "green",
		brightgreen: "brightgreen",
	}
)

func intToColor(value int) string {
	if value < red {
		return colors[-1]
	}
	if value > brightgreen {
		return colors[brightgreen]
	}
	return colors[value]
}

func main() {
	flag.Parse()
	covValue := *covPtr
	if covValue < 0 || covValue > 100 {
		covValue = -1
	} else {
		covValue = covValue * (brightgreen + 1) / 100
	}

	color := intToColor(covValue)
	var shieldsURL string
	if color == "unknown" {
		shieldsURL = "coverage-unknown-lightgrey"
	} else {
		shieldsURL = url.QueryEscape(fmt.Sprintf("coverage-%v%%-%s", *covPtr, color))
	}
	fmt.Print(baseURL + shieldsURL)
}

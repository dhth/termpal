package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	foreground = flag.String("foreground", "#282828", "hex color to use for foreground")
	colorsFl   = flag.String("colors", "", "comma separated list of hex colors")
)

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

const (
	colorsPerRow = 8
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\nFlags:\n", "termpal shows you how colors will look in your terminal.\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	hexRegex := regexp.MustCompile(`^#([A-Fa-f0-9]{3}|[A-Fa-f0-9]{6})$`)

	if !hexRegex.MatchString(*foreground) {
		die("foreground is not valid regex")
	}
	var colors []string

	if *colorsFl == "" {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			c := scanner.Text()
			if hexRegex.MatchString(c) {
				colors = append(colors, scanner.Text())
			}
		}
	} else {
		colorsEls := strings.Split(*colorsFl, ",")
		for _, col := range colorsEls {
			colSt := strings.TrimSpace(col)
			if hexRegex.MatchString(colSt) {
				colors = append(colors, colSt)
			}
		}
	}

	var bgStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(*foreground)).
		PaddingLeft(1).
		PaddingRight(1)

	var fgStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1)

	var counter int

	fmt.Printf("\n")

	for _, color := range colors {
		fmt.Printf("%s\t",
			bgStyle.Copy().Background(lipgloss.Color(color)).Render(color),
		)
		counter++
		if counter >= colorsPerRow {
			fmt.Println()
			counter = 0
		}
	}

	fmt.Printf("\n\n\n")

	counter = 0
	for _, color := range colors {
		fmt.Printf("%s\t",
			fgStyle.Copy().Foreground(lipgloss.Color(color)).Render(color),
		)
		counter++
		if counter >= colorsPerRow {
			fmt.Println()
			counter = 0
		}
	}
	fmt.Printf("\n")
}

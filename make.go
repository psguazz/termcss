package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	css := termcss()

	err := os.WriteFile("term.css", []byte(css), 0o644)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
		return
	}

	fmt.Println("Done!")
}

var palette = map[string]string{
	"bg0":         "#2c2e34",
	"bg1":         "#33353f",
	"bg2":         "#363944",
	"bg3":         "#3b3e48",
	"bg4":         "#414550",
	"bg-blue":     "#85d3f2",
	"bg-dim":      "#222327",
	"bg-green":    "#a7df78",
	"bg-red":      "#ff6077",
	"black":       "#181819",
	"blue":        "#76cce0",
	"diff-blue":   "#354157",
	"diff-green":  "#394634",
	"diff-red":    "#55393d",
	"diff-yellow": "#4e432f",
	"fg":          "#e2e2e3",
	"green":       "#9ed072",
	"grey":        "#7f8490",
	"grey-dim":    "#595f6f",
	"orange":      "#f39660",
	"purple":      "#b39df3",
	"red":         "#fc5d7c",
	"yellow":      "#e7c664",
}

var sides = map[string]string{
	"t": "top",
	"b": "bottom",
	"l": "left",
	"r": "right",
}

var combos = map[string]string{
	"t": "y",
	"b": "y",
	"l": "x",
	"r": "x",
}

var dims = map[string]string{
	"x": "col",
	"y": "row",
}

var (
	positiveSpaces = []int{0, 1, 2}
	negativeSpaces = []int{-2, -1}
)

func declaration(property, value string) string {
	return fmt.Sprintf("%s: %s;", property, value)
}

func rule(selector string, declarations []string) string {
	block := strings.Join(declarations, "\n  ")
	return fmt.Sprintf("%s {\n  %s\n}\n\n", selector, block)
}

func termcss() string {
	var css strings.Builder

	css.WriteString(variables())
	css.WriteString(foundation())
	css.WriteString(typography())
	css.WriteString(colors())
	css.WriteString(flex())
	css.WriteString(spacing())
	css.WriteString(border())
	css.WriteString(sizes())

	return css.String()
}

func variables() string {
	var css strings.Builder

	var colors []string
	for name, hex := range palette {
		colorVar := declaration("--"+name, hex)
		colors = append(colors, colorVar)
	}

	css.WriteString(rule("html", colors))

	css.WriteString(rule("html", []string{
		declaration("--base-size", "13px"),
		declaration("--col", "1rch"),
		declaration("--row", "1rlh"),
	}))

	css.WriteString(rule("html", []string{
		declaration("font-size", "var(--base-size)"),
		declaration("line-height", "var(--row)"),
	}))

	return css.String()
}

func foundation() string {
	var css strings.Builder

	css.WriteString(rule("*, *::before, *::after", []string{
		declaration("padding", "0"),
		declaration("margin", "0"),
		declaration("font-family", "inherit"),
		declaration("box-sizing", "border-box"),
	}))

	css.WriteString(rule("body", []string{
		declaration("font-family", "\"Monaco Nerd Font\""),
		declaration("background-color", "var(--bg-dim)"),
		declaration("color", "var(--fg)"),
		declaration("max-width", "calc(var(--col) * 90)"),
	}))

	css.WriteString(rule("ul", []string{
		declaration("list-style", "none"),
		declaration("padding-left", "calc(var(--col) * 2)"),
		declaration("margin", "0"),
	}))

	css.WriteString(rule("li", []string{
		declaration("position", "relative"),
	}))

	css.WriteString(rule("li:before", []string{
		declaration("content", "\"*\""),
		declaration("position", "absolute"),
		declaration("left", "calc(var(--col) * -2)"),
		declaration("color", "var(--grey)"),
	}))

	return css.String()
}

func typography() string {
	var css strings.Builder

	sizes := []int{1, 2, 3}

	css.WriteString(rule(".text-0\\.5", []string{
		declaration("font-size", "calc(var(--base-size) * 0.5)"),
		declaration("line-height", "calc(var(--row) * 1)"),
	}))

	for _, size := range sizes {
		selector := fmt.Sprintf(".text-%d", size)

		css.WriteString(rule(selector, []string{
			declaration("font-size", fmt.Sprintf("calc(var(--base-size) * %d)", size)),
			declaration("line-height", fmt.Sprintf("calc(var(--row) * %d)", size)),
		}))

	}

	css.WriteString(rule(".font-bold", []string{
		declaration("font-weight", "bold"),
	}))
	css.WriteString(rule(".font-normal", []string{
		declaration("font-weight", "normal"),
	}))

	css.WriteString(rule(".italic", []string{
		declaration("font-style", "italic"),
	}))
	css.WriteString(rule(".not-italic", []string{
		declaration("font-style", "normal"),
	}))

	return css.String()
}

func colors() string {
	var css strings.Builder

	for name := range palette {
		css.WriteString(rule(fmt.Sprintf(".text-%s", name), []string{
			declaration("color", fmt.Sprintf("var(--%s)", name)),
		}))

		css.WriteString(rule(fmt.Sprintf(".bg-%s", name), []string{
			declaration("background-color", fmt.Sprintf("var(--%s)", name)),
		}))

	}

	return css.String()
}

func flex() string {
	var css strings.Builder

	css.WriteString(rule(".flex-row", []string{
		declaration("display", "flex"),
		declaration("flex-direction", "row"),
	}))

	css.WriteString(rule(".flex-col", []string{
		declaration("display", "flex"),
		declaration("flex-direction", "column"),
	}))

	css.WriteString(rule(".grow", []string{
		declaration("flex-grow", "1"),
	}))

	return css.String()
}

func spacing() string {
	var css strings.Builder

	spaces := map[string]string{
		"p": "padding",
		"m": "margin",
	}

	sizes := map[string][]int{
		"p": positiveSpaces,
		"m": append(negativeSpaces, positiveSpaces...),
	}

	for selName, propName := range spaces {
		for _, size := range sizes[selName] {
			for selSide, propSide := range sides {
				combo := combos[selSide]
				sign := ""
				number := size
				if size < 0 {
					sign = "-"
					number = -size
				}

				selBase := fmt.Sprintf(".%s-%d", sign+selName+selSide, number)
				selCombo := fmt.Sprintf(".%s-%d", sign+selName+combo, number)
				selFull := fmt.Sprintf(".%s-%d", sign+selName, number)

				selector := fmt.Sprintf("%s, %s, %s", selBase, selCombo, selFull)
				property := fmt.Sprintf("%s-%s", propName, propSide)
				value := fmt.Sprintf("calc(var(--%s) * %d)", dims[combo], size)

				css.WriteString(rule(selector, []string{
					declaration(property, value),
				}))

			}
		}
	}

	return css.String()
}

func border() string {
	var css strings.Builder
	return css.String()
}

func sizes() string {
	var css strings.Builder
	return css.String()
}

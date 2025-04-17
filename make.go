package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	css := termcss()

	err := os.WriteFile("layout.css", []byte(css), 0o644)
	if err != nil {
		fmt.Println("❌ Failed to write layout.css:", err)
		return
	}

	fmt.Println("✅ layout.css written successfully.")
}

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
	css.WriteString(layout())
	css.WriteString(spacing())
	css.WriteString(border())
	css.WriteString(sizes())

	return css.String()
}

func variables() string {
	var css strings.Builder

	css.WriteString(rule("html", []string{
		declaration("--bg0", "#2c2e34"),
		declaration("--bg1", "#33353f"),
		declaration("--bg2", "#363944"),
		declaration("--bg3", "#3b3e48"),
		declaration("--bg4", "#414550"),
		declaration("--bg-blue", "#85d3f2"),
		declaration("--bg-dim", "#222327"),
		declaration("--bg-green", "#a7df78"),
		declaration("--bg-red", "#ff6077"),
		declaration("--black", "#181819"),
		declaration("--blue", "#76cce0"),
		declaration("--diff-blue", "#354157"),
		declaration("--diff-green", "#394634"),
		declaration("--diff-red", "#55393d"),
		declaration("--diff-yellow", "#4e432f"),
		declaration("--fg", "#e2e2e3"),
		declaration("--green", "#9ed072"),
		declaration("--grey", "#7f8490"),
		declaration("--grey-dim", "#595f6f"),
		declaration("--orange", "#f39660"),
		declaration("--purple", "#b39df3"),
		declaration("--red", "#fc5d7c"),
		declaration("--yellow", "#e7c664"),
	}))

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
	return css.String()
}

func colors() string {
	var css strings.Builder

	cols := []string{
		"bg0",
		"bg1",
		"bg2",
		"bg3",
		"bg4",
		"bg-blue",
		"bg-dim",
		"bg-green",
		"bg-red",
		"black",
		"blue",
		"diff-blue",
		"diff-green",
		"diff-red",
		"diff-yellow",
		"fg",
		"green",
		"grey",
		"grey-dim",
		"orange",
		"purple",
		"red",
		"yellow",
	}

	for _, color := range cols {
		css.WriteString(rule(fmt.Sprintf(".text-%s", color), []string{
			declaration("color", fmt.Sprintf("var(--%s)", color)),
		}))

		css.WriteString(rule(fmt.Sprintf(".bg-%s", color), []string{
			declaration("background-color", fmt.Sprintf("var(--%s)", color)),
		}))

	}

	return css.String()
}

func layout() string {
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
		"p": {0, 1, 2},
		"m": {-2, -1, 0, 1, 2},
	}

	sides := map[string]string{
		"t": "top",
		"b": "bottom",
		"l": "left",
		"r": "right",
	}

	combos := map[string]string{
		"t": "y",
		"b": "y",
		"l": "x",
		"r": "x",
	}

	dims := map[string]string{
		"x": "col",
		"y": "row",
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

				selBase := fmt.Sprintf("%s-%d", sign+selName+selSide, number)
				selCombo := fmt.Sprintf("%s-%d", sign+selName+combo, number)
				selFull := fmt.Sprintf("%s-%d", sign+selName, number)

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

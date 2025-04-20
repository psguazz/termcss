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
	"transparent": "transparent",
	"inherit":     "inherit",
}

var axes = []string{"x", "y"}

var sides = map[string][]string{
	"":  {"t", "b", "l", "r"},
	"x": {"l", "r"},
	"y": {"t", "b"},
	"t": {"t"},
	"b": {"b"},
	"l": {"l"},
	"r": {"r"},
}

var corners = map[string][]string{
	"":   {"tl", "tr", "bl", "br"},
	"x":  {"tl", "tr", "bl", "br"},
	"y":  {"tl", "tr", "bl", "br"},
	"t":  {"tl", "tr"},
	"b":  {"bl", "br"},
	"l":  {"tl", "bl"},
	"r":  {"tr", "br"},
	"tl": {"tl"},
	"tr": {"tr"},
	"bl": {"bl"},
	"br": {"br"},
}

var (
	positiveSpaces = []int{0, 1, 2, 4}
	negativeSpaces = []int{-2, -1}
)

var (
	sizes   = []int{0, 1, 2, 4, 8, 12, 20, 32, 40, 80, 100, 120}
	offsets = []int{0, 1, 2, 3, 4, 6, 8, 12, 20}
)

var props = map[string]string{
	"x":  "x",
	"y":  "y",
	"t":  "top",
	"b":  "bottom",
	"l":  "left",
	"r":  "right",
	"tl": "top-left",
	"tr": "top-right",
	"bl": "bottom-left",
	"br": "bottom-right",
}

var dims = map[string]string{
	"x": "col",
	"y": "row",
	"t": "row",
	"b": "row",
	"l": "col",
	"r": "col",
	"h": "row",
	"w": "col",
}

func grid(dim string, n int) string {
	return fmt.Sprintf("calc(var(--%s) * %d)", dim, n)
}

func declaration(property, value string) string {
	return fmt.Sprintf("%s: %s;", property, value)
}

func rule(selector string, declarations []string) string {
	block := strings.Join(declarations, "\n  ")
	return fmt.Sprintf("%s {\n  %s\n}\n\n", selector, block)
}

func docs(section string, lines []string) string {
	title := fmt.Sprintf("/* %s\n * ====================", strings.ToUpper(section))
	description := strings.Join(lines, "\n * ")
	return fmt.Sprintf("%s\n * %s\n */\n\n", title, description)
}

func stringMap(input []string, f func(string) string) []string {
	output := make([]string, len(input))
	for i, v := range input {
		output[i] = f(v)
	}
	return output
}

func termcss() string {
	var css strings.Builder

	css.WriteString(variables())
	css.WriteString(foundation())
	css.WriteString(typography())
	css.WriteString(colors())
	css.WriteString(flex())
	css.WriteString(spacing())
	css.WriteString(borders())
	css.WriteString(sizing())
	css.WriteString(positioning())

	return css.String()
}

func variables() string {
	var css strings.Builder

	var colors []string
	for name, hex := range palette {
		colorVar := declaration("--"+name, hex)
		colors = append(colors, colorVar)
	}

	css.WriteString(docs("Variables", []string{
		"Setting up some variables used throughout the styles:",
		"- All the colors. These come from Sonokai because it's awesome.",
		"- Font size, line height, row width. The Grid.",
	}))

	css.WriteString(rule("html", colors))

	css.WriteString(rule("html", []string{
		declaration("--base-size", "14px"),
		declaration("--col", "calc(1rem * 0.6007142857)"),
		declaration("--row", "calc(1rem * 1.35)"),
	}))

	css.WriteString(rule("html", []string{
		declaration("font-size", "var(--base-size)"),
		declaration("line-height", "var(--row)"),
	}))

	return css.String()
}

func foundation() string {
	var css strings.Builder

	css.WriteString(docs("Foundation", []string{
		"Some default properties that allow us to apply the theme and guarantee grid alignment.",
		"- The list elements are a bit opinionated.",
		"- We _need_ `content-box` to enforce the grid when applying widths and heights to items with borders",
	}))

	css.WriteString(rule("*, *::before, *::after", []string{
		declaration("padding", "0"),
		declaration("margin", "0"),
		declaration("border", "0"),
		declaration("font-size", "inherit"),
		declaration("color", "inherit"),
		declaration("background-color", "none"),
		declaration("font-family", "inherit"),
		declaration("box-sizing", "content-box"),
	}))

	css.WriteString(rule("body", []string{
		declaration("font-family", "\"Monaco Nerd Font\""),
		declaration("background-color", "var(--bg-dim)"),
		declaration("color", "var(--fg)"),
	}))

	css.WriteString(rule("ul", []string{
		declaration("list-style", "none"),
		declaration("padding-left", grid("col", 2)),
		declaration("margin", "0"),
	}))

	css.WriteString(rule("li", []string{
		declaration("position", "relative"),
	}))

	css.WriteString(rule("li:before", []string{
		declaration("content", "\"*\""),
		declaration("position", "absolute"),
		declaration("left", grid("col", -2)),
		declaration("color", "var(--grey)"),
	}))

	return css.String()
}

func typography() string {
	var css strings.Builder

	sizes := []int{1, 2, 3}

	css.WriteString(docs("Typography", []string{
		"Basic text transofrmations: weight, style, whitespace, case.",
		"Size is also an option. It's definitely cheating a bit, since terminals don't have that luxury,",
		"but it's nice to have and is still guaranteed to be aligned.",
	}))

	for _, size := range sizes {
		selector := fmt.Sprintf(".text-%d", size)

		css.WriteString(rule(selector, []string{
			declaration("font-size", grid("base-size", size)),
			declaration("line-height", grid("row", size)),
		}))

	}

	whitespaces := []string{"normal", "nowrap"}
	for _, whitespace := range whitespaces {
		css.WriteString(rule(".whitespace-"+whitespace, []string{
			declaration("white-space", whitespace),
		}))
	}

	weights := []string{"normal", "bold"}
	for _, weight := range weights {
		css.WriteString(rule(".font-"+weight, []string{
			declaration("font-weight", weight),
		}))
	}

	transforms := []string{"uppercase", "lowercase", "capitalize"}
	for _, transform := range transforms {
		css.WriteString(rule("."+transform, []string{
			declaration("text-transform", transform),
		}))
	}

	css.WriteString(rule(".normal-case", []string{
		declaration("text-transform", "none"),
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

	css.WriteString(docs("Colors", []string{
		"Utilities to apply the theme colors to text and background.",
	}))

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

	css.WriteString(docs("Flex", []string{
		"Flex-related utilities. Still WIP.",
		"Note that some things like `justify-between will not be available unless they can be guaranteed to _always_ respect the grid.",
	}))

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

	css.WriteString(rule(".shrink", []string{
		declaration("flex-shrink", "1"),
	}))

	return css.String()
}

func spacing() string {
	var css strings.Builder

	css.WriteString(docs("Margin & Padding", []string{
		"Applying margin and padding -- specific sides, x and y, all sides.",
		"All the sizes are grid-aligned.",
		"Margins are also negative.",
	}))

	spaces := map[string]string{
		"p": "padding",
		"m": "margin",
	}

	sizes := map[string][]int{
		"p": positiveSpaces,
		"m": append(negativeSpaces, positiveSpaces...),
	}

	addSpacing := func(selName string, baseMod string, size int) {
		selector := selName + baseMod
		if size > 0 {
			selector = fmt.Sprintf(".%s-%d", selector, size)
		} else {
			selector = fmt.Sprintf(".-%s-%d", selector, -size)
		}

		css.WriteString(rule(selector, stringMap(sides[baseMod], func(mod string) string {
			property := fmt.Sprintf("%s-%s", spaces[selName], props[mod])
			return declaration(property, grid(dims[mod], size))
		})))
	}

	for selName := range spaces {
		for _, size := range sizes[selName] {
			addSpacing(selName, "", size)
		}

		for _, axis := range axes {
			for _, size := range sizes[selName] {
				addSpacing(selName, axis, size)
			}

			for _, side := range sides[axis] {
				for _, size := range sizes[selName] {
					addSpacing(selName, side, size)
				}
			}
		}
	}

	css.WriteString(docs("Margin Auto", []string{
		"This is incuded for the sake of creating decent page layouts.",
		"It's fine to use this on the root element, but anywhere else",
		"it might break the grid.",
	}))

	css.WriteString(rule(".m-auto", []string{
		declaration("margin", "auto"),
	}))
	css.WriteString(rule(".mx-auto", []string{
		declaration("margin-inline", "auto"),
	}))
	css.WriteString(rule(".my-auto", []string{
		declaration("margin-block", "auto"),
	}))

	return css.String()
}

func borders() string {
	var css strings.Builder

	css.WriteString(docs("Borders", []string{
		"Applying borders -- specific sides, x and y, all sides.",
		"Colors, style, and radius can be changed.",
		"The border width cannot be chosen; it's just yes or no.",
		"Borders are applied along with a specific combination of margin and padding so that",
		"they take a full extra row/column and remain aligned grid.",
		"These half-paddings are the reason we need `content-box`.",
		"Borders are defined _after_ margins and paddings, meaning that a bordered item cannot have additional margin or padding.",
		"This is, again, to enforce grid alignment at all times",
	}))

	borderWidth := 2
	styles := []string{"solid", "dashed"}
	radiuses := map[string]string{"square": "0", "rounded": "4px"}

	halfRow := fmt.Sprintf("calc((var(--%s) - %dpx) / 2)", "row", borderWidth)
	halfCol := fmt.Sprintf("calc((var(--%s) - %dpx) / 2)", "col", borderWidth)
	border := fmt.Sprintf("%dpx solid var(--grey-dim)", borderWidth)

	halfDims := map[string]string{
		"x": halfCol,
		"y": halfRow,
		"t": halfRow,
		"b": halfRow,
		"l": halfCol,
		"r": halfCol,
	}

	addBorders := func(baseMod string) {
		baseSel := ".border"
		if baseMod != "" {
			baseSel = fmt.Sprintf("%s-%s", baseSel, baseMod)
		}

		borderDecl := stringMap(sides[baseMod], func(mod string) string {
			property := fmt.Sprintf("border-%s", props[mod])
			return declaration(property, border)
		})

		marginDecl := stringMap(sides[baseMod], func(mod string) string {
			property := fmt.Sprintf("margin-%s", props[mod])
			return declaration(property, halfDims[mod])
		})

		paddingDecl := stringMap(sides[baseMod], func(mod string) string {
			property := fmt.Sprintf("padding-%s", props[mod])
			return declaration(property, halfDims[mod])
		})

		css.WriteString(rule(baseSel, append(append(borderDecl, marginDecl...), paddingDecl...)))

		borderDeclNo := stringMap(sides[baseMod], func(mod string) string {
			property := fmt.Sprintf("border-%s", props[mod])
			return declaration(property, "0")
		})

		marginDeclNo := stringMap(sides[baseMod], func(mod string) string {
			property := fmt.Sprintf("margin-%s", props[mod])
			return declaration(property, "0")
		})

		paddingDeclNo := stringMap(sides[baseMod], func(mod string) string {
			property := fmt.Sprintf("padding-%s", props[mod])
			return declaration(property, "0")
		})

		css.WriteString(rule(baseSel+"-0", append(append(borderDeclNo, marginDeclNo...), paddingDeclNo...)))

		for color := range palette {
			selector := fmt.Sprintf("%s-%s", baseSel, color)
			value := fmt.Sprintf("var(--%s)", color)

			css.WriteString(rule(selector, stringMap(sides[baseMod], func(mod string) string {
				property := fmt.Sprintf("border-%s-color", props[mod])
				return declaration(property, value)
			})))
		}

		for _, style := range styles {
			selector := fmt.Sprintf("%s-%s", baseSel, style)

			css.WriteString(rule(selector, stringMap(sides[baseMod], func(mod string) string {
				property := fmt.Sprintf("border-%s-style", props[mod])
				return declaration(property, style)
			})))
		}
	}

	addCorners := func(baseMod string) {
		baseSel := ".border"
		if baseMod != "" {
			baseSel = fmt.Sprintf("%s-%s", baseSel, baseMod)
		}

		for selMod, radius := range radiuses {
			selector := fmt.Sprintf("%s-%s", baseSel, selMod)

			css.WriteString(rule(selector, stringMap(corners[baseMod], func(mod string) string {
				property := fmt.Sprintf("border-%s-radius", props[mod])
				return declaration(property, radius)
			})))
		}
	}

	addBorders("")
	addCorners("")

	for _, axis := range axes {
		addBorders(axis)
		addCorners(axis)

		for _, side := range sides[axis] {
			addBorders(side)
			addCorners(side)

			for _, corner := range corners[side] {
				addCorners(corner)
			}
		}
	}

	css.WriteString(docs("Merged Borders", []string{
		"These utilities are essentially the eqivalent of `-m-1` to use alongside border classes.",
		"They are offered as convenience classes to combine containers with adjacent borders.",
		"The alternative to this would be to nest the bordered containers in something else,",
		"and apply the negative margins to the outer elements; but this can have bad side-effects",
		"in more complex layouts, in addition to just complicate things that should be simple.",
		"WARNING: These classes will break the grid unless used together with the corresponding borders",
	}))

	addMerges := func(baseMod string) {
		selector := ".border-merge"
		if baseMod != "" {
			selector = fmt.Sprintf("%s-%s", selector, baseMod)
		}

		css.WriteString(rule(selector, stringMap(sides[baseMod], func(mod string) string {
			property := "margin-" + props[mod]
			value := fmt.Sprintf("calc((var(--%s) + %dpx) / -2)", dims[mod], borderWidth)
			return declaration(property, value)
		})))
	}

	addMerges("")

	for _, axis := range axes {
		addMerges(axis)

		for _, side := range sides[axis] {
			addMerges(side)
		}
	}

	return css.String()
}

func sizing() string {
	var css strings.Builder

	css.WriteString(docs("Widths & Heights", []string{
		"A collection of width and height utilities.",
		"Each option comes with a `min-` and `max-` variant.",
	}))

	rules := map[string]string{
		"w": "width",
		"h": "height",
	}

	mods := []string{"", "max-", "min-"}

	for sel, property := range rules {
		for _, mod := range mods {
			for _, size := range sizes {
				selector := fmt.Sprintf(".%s-%d", mod+sel, size)
				value := grid(dims[sel], size)

				css.WriteString(rule(selector, []string{
					declaration(mod+property, value),
				}))
			}
		}
	}

	return css.String()
}

func positioning() string {
	var css strings.Builder

	css.WriteString(docs("Positioning", []string{
		"Position, display, and visibility (including z-index).",
		"Also including some `top`, `bottom`, `left`, and `right` options.",
	}))

	positions := []string{"static", "fixed", "absolute", "relative", "sticky"}
	for _, position := range positions {
		css.WriteString(rule("."+position, []string{
			declaration("position", position),
		}))
	}

	displays := []string{"inline", "block", "inline-block"}
	for _, display := range displays {
		css.WriteString(rule("."+display, []string{
			declaration("display", display),
		}))
	}

	for _, axis := range axes {
		for _, side := range sides[axis] {
			for _, offset := range offsets {
				selector := fmt.Sprintf(".%s-%d", side, offset)
				value := grid(dims[side], offset)

				css.WriteString(rule(selector, []string{
					declaration(props[side], value),
				}))
			}
		}
	}

	visibilities := []string{"visible", "hidden"}
	for _, visibility := range visibilities {
		css.WriteString(rule("."+visibility, []string{
			declaration("visibility", visibility),
		}))
	}

	overflows := []string{"auto", "hidden", "visible", "scroll"}
	for _, overflow := range overflows {
		css.WriteString(rule(".overflow-"+overflow, []string{
			declaration("overflow", overflow),
		}))
	}

	zIndexes := []string{"1", "10", "100", "1000"}
	for _, zIndex := range zIndexes {
		css.WriteString(rule(".z-"+zIndex, []string{
			declaration("z-index", zIndex),
		}))
	}

	return css.String()
}

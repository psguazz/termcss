<div align="center">
  <img src="logo.png" alt="logo" height="100px">
  <div><strong>Downgrading the web to something better™</strong></div>
</div>

<br>

Lots of people are trying to make the terminal look more like a real GUI, but
what people really want is GUIs that look like a terminal.

Enter `termcss`: a modern, vintage CSS framework that aims at bridging the gap
between the terminal and the web by making the web worse.

## Philosophy

### Utility first

`termcss` follows the footsteps of [Tailwind](https://tailwindcss.com/) and
exposes a series of utility classes to style your documents. To a large extent,
these classes use the same names as Tailwind, the main difference being that
there are _a lot_ fewer.

### The Terminal Grid

The golden rule of this framework is that **everything must line up with the
terminal grid at all times, at all costs.**

The most important aspects of this are:

- A monospace font
- The `--row` and `--col` CSS variables that specify the dimensions of a
  character. In a perfect world they would be set to `1lh` and `1ch`, but these
  units are _not_ reliable enough.

Everything else that affects sizes and positions and all that is defined as
multiple of these variables. As long as you stick to the classes offered here,
your text should never, ever go out of whack.

## Features and examples

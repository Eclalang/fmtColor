package fmtColor

import (
	"fmt"
	"image/color"
)

// PrintRGB prints the values in the given color
func PrintRGB(r uint8, g uint8, b uint8, values ...any) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprint(values...))
}

// PrintlnRGB prints the values in the given color and adds a newline
func PrintlnRGB(r uint8, g uint8, b uint8, values ...any) {
	PrintRGB(r, g, b, values...)
	fmt.Println()
}

// PrintfRGB prints the values in the given color and formats them
func PrintfRGB(r uint8, g uint8, b uint8, toFormat string, values ...any) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprintf(toFormat, values...))
}

// PrintflnRGB prints the values in the given color, formats them and adds a newline
func PrintflnRGB(r uint8, g uint8, b uint8, toFormat string, values ...any) {
	PrintfRGB(r, g, b, toFormat, values...)
	fmt.Println()
}

// Print prints the values in the given color
func Print(c color.Color, values ...any) {
	r, g, b, _ := c.RGBA()
	PrintRGB(uint8(r), uint8(g), uint8(b), values...)
}

// Println prints the values in the given color and adds a newline
func Println(c color.Color, values ...any) {
	r, g, b, _ := c.RGBA()
	PrintlnRGB(uint8(r), uint8(g), uint8(b), values...)
}

// Printf prints the values in the given color and formats them
func Printf(c color.Color, format string, values ...any) {
	r, g, b, _ := c.RGBA()
	PrintfRGB(uint8(r), uint8(g), uint8(b), format, values...)
}

// Printfln prints the values in the given color, formats them and adds a newline
func Printfln(c color.Color, format string, values ...any) {
	r, g, b, _ := c.RGBA()
	PrintflnRGB(uint8(r), uint8(g), uint8(b), format, values...)
}

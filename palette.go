// palette.go
package main

import "fmt"

type Color struct {
	R, G, B uint8
}

func (c Color) ANSI() string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

// Reset is the ANSI reset sequence.
const Reset = "\033[0m"

// Palette is the Dracula ANSI color palette (excluding black and white).
var Palette = []Color{
	{255, 85, 85},   // AnsiRed           #FF5555
	{255, 110, 110}, // AnsiBrightRed     #FF6E6E
	{255, 121, 198}, // AnsiMagenta       #FF79C6
	{255, 146, 223}, // AnsiBrightMagenta #FF92DF
	{189, 147, 249}, // AnsiBlue          #BD93F9
	{214, 172, 255}, // AnsiBrightBlue    #D6ACFF
	{98, 114, 164},  // AnsiBrightBlack   #6272A4
	{139, 233, 253}, // AnsiCyan          #8BE9FD
	{164, 255, 255}, // AnsiBrightCyan    #A4FFFF
	{80, 250, 123},  // AnsiGreen         #50FA7B
	{105, 255, 148}, // AnsiBrightGreen   #69FF94
	{241, 250, 140}, // AnsiYellow        #F1FA8C
	{255, 255, 165}, // AnsiBrightYellow  #FFFFA5
}

// Wraps with correct modulo for negative indices.
func ColorAt(i int) Color {
	n := len(Palette)
	return Palette[((i%n)+n)%n]
}

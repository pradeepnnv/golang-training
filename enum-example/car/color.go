package car

import (
	"strings"
)

//go:generate stringer -type=Color

type Color int

const (
	UnknownColor Color = iota
	Gray
	Green
	Yellow
	White
	Red
)

func (c Color) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *Color) UnmarshalText(text []byte) error {
	*c = ColorFromText(string(text))
	return nil
}

func ColorFromText(c string) Color {
	switch strings.ToLower(c) {
	default:
		return UnknownColor
	case "gray":
		return Gray
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "white":
		return White
	case "red":
		return Red
	}
}

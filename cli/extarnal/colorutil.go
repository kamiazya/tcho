package extarnal

import (
	"github.com/kamiazya/tcho/core/domain/value/colorcode"
	"github.com/fatih/color"
)

type TechoCliColor struct{}

func (tlc TechoCliColor) ColorCode(colorStr string) colorcode.ColorCode {
	switch colorStr {
	case "red":
		return colorcode.Red
	case "green":
		return colorcode.Green
	case "yellow":
		return colorcode.Yellow
	case "blue":
		return colorcode.Blue
	case "magenta":
		return colorcode.Magenta
	case "cyan":
		return colorcode.Cyan
	}
	return colorcode.NoColor
}

func (tlc TechoCliColor) Color(code colorcode.ColorCode) *color.Color {
	switch code {
	case colorcode.Red:
		return color.New(color.FgRed)
	case colorcode.Green:
		return color.New(color.FgGreen)
	case colorcode.Yellow:
		return color.New(color.FgYellow)
	case colorcode.Blue:
		return color.New(color.FgBlue)
	case colorcode.Magenta:
		return color.New(color.FgMagenta)
	case colorcode.Cyan:
		return color.New(color.FgCyan)
	}
	return color.New()
}

func (tlc TechoCliColor) ColorBg(code colorcode.ColorCode) *color.Color {
	switch code {
	case colorcode.Red:
		return color.New(color.BgRed)
	case colorcode.Green:
		return color.New(color.BgGreen)
	case colorcode.Yellow:
		return color.New(color.BgYellow)
	case colorcode.Blue:
		return color.New(color.BgBlue)
	case colorcode.Magenta:
		return color.New(color.BgMagenta)
	case colorcode.Cyan:
		return color.New(color.BgCyan)
	}
	return color.New()
}

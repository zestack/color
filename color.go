package color

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	output    io.Writer
	isColored bool
)

func init() {
	SetOutput(os.Stderr)
}

func SetOutput(out io.Writer) {
	output = out
	if w, ok := out.(Writer); ok {
		isColored = w.Colorable()
	} else if out == os.Stdout || out == os.Stderr {
		isColored = os.Getenv("NO_COLOR") == "" &&
			os.Getenv("TERM") != "dumb" &&
			Colorable(out)
	} else {
		isColored = Colorable(out)
	}
}

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

const (
	ResetBold Attribute = iota + 22
	ResetItalic
	ResetUnderline
	ResetBlinking
	_
	ResetReversed
	ResetConcealed
	ResetCrossedOut
)

var mapResetAttributes map[Attribute]Attribute = map[Attribute]Attribute{
	Bold:         ResetBold,
	Faint:        ResetBold,
	Italic:       ResetItalic,
	Underline:    ResetUnderline,
	BlinkSlow:    ResetBlinking,
	BlinkRapid:   ResetBlinking,
	ReverseVideo: ResetReversed,
	Concealed:    ResetConcealed,
	CrossedOut:   ResetCrossedOut,
}

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

type Color struct {
	attrs    []Attribute
	disabled *bool
}

func New(attrs ...Attribute) *Color {
	return &Color{attrs, nil}
}

func (c *Color) Enable(v ...bool) *Color {
	if len(v) > 0 {
		c.disabled = boolPtr(v[0] == false)
	} else {
		c.disabled = boolPtr(false)
	}
	return c
}

func (c *Color) Enabled() bool {
	if c.disabled == nil {
		return true
	}
	return !*c.disabled
}

func (c *Color) Add(attrs ...Attribute) *Color {
	c.attrs = append(c.attrs, attrs...)
	return c
}

func (c *Color) FgBlack() *Color      { return c.Add(FgBlack) }
func (c *Color) FgRed() *Color        { return c.Add(FgRed) }
func (c *Color) FgGreen() *Color      { return c.Add(FgGreen) }
func (c *Color) FgYellow() *Color     { return c.Add(FgYellow) }
func (c *Color) FgBlue() *Color       { return c.Add(FgBlue) }
func (c *Color) FgMagenta() *Color    { return c.Add(FgMagenta) }
func (c *Color) FgCyan() *Color       { return c.Add(FgCyan) }
func (c *Color) FgWhite() *Color      { return c.Add(FgWhite) }
func (c *Color) FgHiBlack() *Color    { return c.Add(FgHiBlack) }
func (c *Color) FgHiRed() *Color      { return c.Add(FgHiRed) }
func (c *Color) FgHiGreen() *Color    { return c.Add(FgHiGreen) }
func (c *Color) FgHiYellow() *Color   { return c.Add(FgHiYellow) }
func (c *Color) FgHiBlue() *Color     { return c.Add(FgHiBlue) }
func (c *Color) FgHiMagenta() *Color  { return c.Add(FgHiMagenta) }
func (c *Color) FgHiCyan() *Color     { return c.Add(FgHiCyan) }
func (c *Color) FgHiWhite() *Color    { return c.Add(FgHiWhite) }
func (c *Color) BgBlack() *Color      { return c.Add(BgBlack) }
func (c *Color) BgRed() *Color        { return c.Add(BgRed) }
func (c *Color) BgGreen() *Color      { return c.Add(BgGreen) }
func (c *Color) BgYellow() *Color     { return c.Add(BgYellow) }
func (c *Color) BgBlue() *Color       { return c.Add(BgBlue) }
func (c *Color) BgMagenta() *Color    { return c.Add(BgMagenta) }
func (c *Color) BgCyan() *Color       { return c.Add(BgCyan) }
func (c *Color) BgWhite() *Color      { return c.Add(BgWhite) }
func (c *Color) BgHiBlack() *Color    { return c.Add(BgHiBlack) }
func (c *Color) BgHiRed() *Color      { return c.Add(BgHiRed) }
func (c *Color) BgHiGreen() *Color    { return c.Add(BgHiGreen) }
func (c *Color) BgHiYellow() *Color   { return c.Add(BgHiYellow) }
func (c *Color) BgHiBlue() *Color     { return c.Add(BgHiBlue) }
func (c *Color) BgHiMagenta() *Color  { return c.Add(BgHiMagenta) }
func (c *Color) BgHiCyan() *Color     { return c.Add(BgHiCyan) }
func (c *Color) BgHiWhite() *Color    { return c.Add(BgHiWhite) }
func (c *Color) Bold() *Color         { return c.Add(Bold) }
func (c *Color) Faint() *Color        { return c.Add(Faint) }
func (c *Color) Italic() *Color       { return c.Add(Italic) }
func (c *Color) Underline() *Color    { return c.Add(Underline) }
func (c *Color) BlinkSlow() *Color    { return c.Add(BlinkSlow) }
func (c *Color) BlinkRapid() *Color   { return c.Add(BlinkRapid) }
func (c *Color) ReverseVideo() *Color { return c.Add(ReverseVideo) }
func (c *Color) Concealed() *Color    { return c.Add(Concealed) }
func (c *Color) CrossedOut() *Color   { return c.Add(CrossedOut) }

func (c *Color) Wrap(value any) *Value {
	return NewValue(value, c.attrs...)
}

func (c *Color) Bytes() []byte {
	return Bytes(c.attrs...)
}

func (c *Color) Print(a ...any) (n int, err error) {
	c.setWriter(output, isColored)
	defer c.unsetWriter(output, isColored)
	return fmt.Fprint(output, a...)
}

func (c *Color) Printf(format string, a ...any) (n int, err error) {
	c.setWriter(output, isColored)
	defer c.unsetWriter(output, isColored)
	return fmt.Fprintf(output, format, a...)
}

func (c *Color) Println(a ...any) (n int, err error) {
	return fmt.Fprintln(output, c.wrap(sprintln(a...), isColored))
}

func (c *Color) Fprint(w io.Writer, a ...any) (n int, err error) {
	colored := Colorable(w)
	c.setWriter(w, colored)
	defer c.unsetWriter(w, colored)
	return fmt.Fprint(w, a...)
}

func (c *Color) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	colored := Colorable(w)
	c.setWriter(w, colored)
	defer c.unsetWriter(w, colored)
	return fmt.Fprintf(w, format, a...)
}

func (c *Color) Fprintln(w io.Writer, a ...any) (n int, err error) {
	return fmt.Fprintln(w, c.wrap(sprintln(a...), Colorable(w)))
}

func (c *Color) Sprint(a ...any) string {
	return c.wrap(fmt.Sprint(a...), true)
}

func (c *Color) Sprintf(format string, a ...any) string {
	return c.wrap(fmt.Sprintf(format, a...), true)
}

func (c *Color) Sprintln(a ...any) string {
	return c.wrap(sprintln(a...), true) + "\n"
}

// Set sets the SGR sequence.
func (c *Color) setWriter(w io.Writer, colored bool) {
	if c.isColorSet(colored) {
		fmt.Fprint(w, c.format())
	}
}

func (c *Color) unsetWriter(w io.Writer, colored bool) {
	if c.isColorSet(colored) {
		fmt.Fprintf(w, "%s[%dm", escape, Reset)
	}
}

func (c *Color) isColorSet(colored ...bool) bool {
	if c.disabled != nil {
		return !*c.disabled
	} else if len(colored) > 0 {
		return colored[0]
	} else {
		return true
	}
}

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.attrs))
	for i, v := range c.attrs {
		format[i] = strconv.Itoa(int(v))
	}
	return strings.Join(format, ";")
}

func (c *Color) wrap(str string, colored bool) string {
	if c.isColorSet(colored) {
		return c.format() + str + c.unformat()
	}
	return str
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	//return fmt.Sprintf("%s[%dm", escape, Reset)
	//for each element in sequence let's use the speficic reset escape, ou the generic one if not found
	format := make([]string, len(c.attrs))
	for i, v := range c.attrs {
		format[i] = strconv.Itoa(int(Reset))
		ra, ok := mapResetAttributes[v]
		if ok {
			format[i] = strconv.Itoa(int(ra))
		}
	}
	return fmt.Sprintf("%s[%sm", escape, strings.Join(format, ";"))
}

// sprintln is a helper function to format a string with fmt.Sprintln and trim the trailing newline.
func sprintln(a ...any) string {
	return strings.TrimSuffix(fmt.Sprintln(a...), "\n")
}

func boolPtr(v bool) *bool {
	return &v
}

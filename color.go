package color

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

type inner func(any, []string, *Color) string

// Color styles
const (
	// Blk Black text style
	Blk = "30"
	// Rd red text style
	Rd = "31"
	// Grn green text style
	Grn = "32"
	// Yel yellow text style
	Yel = "33"
	// Blu blue text style
	Blu = "34"
	// Mgn magenta text style
	Mgn = "35"
	// Cyn cyan text style
	Cyn = "36"
	// Wht white text style
	Wht = "37"
	// Gry gray text style
	Gry = "90"

	// BlkBg black background style
	BlkBg = "40"
	// RdBg red background style
	RdBg = "41"
	// GrnBg green background style
	GrnBg = "42"
	// YelBg yellow background style
	YelBg = "43"
	// BluBg blue background style
	BluBg = "44"
	// MgnBg magenta background style
	MgnBg = "45"
	// CynBg cyan background style
	CynBg = "46"
	// WhtBg white background style
	WhtBg = "47"

	// R reset emphasis style
	R = "0"
	// B bold emphasis style
	B = "1"
	// D dim emphasis style
	D = "2"
	// I italic emphasis style
	I = "3"
	// U underline emphasis style
	U = "4"
	// In inverse emphasis style
	In = "7"
	// H hidden emphasis style
	H = "8"
	// S strikeout emphasis style
	S = "9"
)

var (
	black   = outer(Blk)
	red     = outer(Rd)
	green   = outer(Grn)
	yellow  = outer(Yel)
	blue    = outer(Blu)
	magenta = outer(Mgn)
	cyan    = outer(Cyn)
	white   = outer(Wht)
	grey    = outer(Gry)

	blackBg   = outer(BlkBg)
	redBg     = outer(RdBg)
	greenBg   = outer(GrnBg)
	yellowBg  = outer(YelBg)
	blueBg    = outer(BluBg)
	magentaBg = outer(MgnBg)
	cyanBg    = outer(CynBg)
	whiteBg   = outer(WhtBg)

	reset     = outer(R)
	bold      = outer(B)
	dim       = outer(D)
	italic    = outer(I)
	underline = outer(U)
	inverse   = outer(In)
	hidden    = outer(H)
	strikeout = outer(S)

	global = New()
)

func outer(n string) inner {
	return func(msg any, styles []string, c *Color) string {
		// TODO: Drop fmt to boost performance?
		if c.disabled {
			return fmt.Sprintf("%v", msg)
		}
		b := new(bytes.Buffer)
		b.WriteString("\x1b[")
		b.WriteString(n)
		for _, s := range styles {
			b.WriteString(";")
			b.WriteString(s)
		}
		b.WriteString("m")
		b.WriteString(fmt.Sprintf("%v", msg))
		b.WriteString("\x1b[0m")
		return b.String()
	}
}

type Color struct {
	output   io.Writer
	disabled bool
}

// New creates a Color instance.
func New() (c *Color) {
	c = new(Color)
	c.SetOutput(colorable.NewColorableStdout())
	return
}

func NewWithOutput(w io.Writer) *Color {
	c := New()
	c.SetOutput(w)
	return c
}

// Output returns the output.
func (c *Color) Output() io.Writer {
	return c.output
}

// SetOutput sets the output.
func (c *Color) SetOutput(w io.Writer) {
	c.output = w
	if w, ok := w.(*os.File); !ok || !isatty.IsTerminal(w.Fd()) {
		c.disabled = true
	}
}

// Disable disables the colors and styles.
func (c *Color) Disable() {
	c.disabled = true
}

// Enable enables the colors and styles.
func (c *Color) Enable() {
	c.disabled = false
}

// Print is analogous to `fmt.Print` with terminal detection.
func (c *Color) Print(args ...any) {
	fmt.Fprint(c.output, args...)
}

// Println is analogous to `fmt.Println` with terminal detection.
func (c *Color) Println(args ...any) {
	fmt.Fprintln(c.output, args...)
}

// Printf is analogous to `fmt.Printf` with terminal detection.
func (c *Color) Printf(format string, args ...any) {
	fmt.Fprintf(c.output, format, args...)
}

func (c *Color) Black(msg any, styles ...string) string {
	return black(msg, styles, c)
}

func (c *Color) Red(msg any, styles ...string) string {
	return red(msg, styles, c)
}

func (c *Color) Green(msg any, styles ...string) string {
	return green(msg, styles, c)
}

func (c *Color) Yellow(msg any, styles ...string) string {
	return yellow(msg, styles, c)
}

func (c *Color) Blue(msg any, styles ...string) string {
	return blue(msg, styles, c)
}

func (c *Color) Magenta(msg any, styles ...string) string {
	return magenta(msg, styles, c)
}

func (c *Color) Cyan(msg any, styles ...string) string {
	return cyan(msg, styles, c)
}

func (c *Color) White(msg any, styles ...string) string {
	return white(msg, styles, c)
}

func (c *Color) Gray(msg any, styles ...string) string {
	return grey(msg, styles, c)
}

func (c *Color) BlackBg(msg any, styles ...string) string {
	return blackBg(msg, styles, c)
}

func (c *Color) RedBg(msg any, styles ...string) string {
	return redBg(msg, styles, c)
}

func (c *Color) GreenBg(msg any, styles ...string) string {
	return greenBg(msg, styles, c)
}

func (c *Color) YellowBg(msg any, styles ...string) string {
	return yellowBg(msg, styles, c)
}

func (c *Color) BlueBg(msg any, styles ...string) string {
	return blueBg(msg, styles, c)
}

func (c *Color) MagentaBg(msg any, styles ...string) string {
	return magentaBg(msg, styles, c)
}

func (c *Color) CyanBg(msg any, styles ...string) string {
	return cyanBg(msg, styles, c)
}

func (c *Color) WhiteBg(msg any, styles ...string) string {
	return whiteBg(msg, styles, c)
}

func (c *Color) Reset(msg any, styles ...string) string {
	return reset(msg, styles, c)
}

func (c *Color) Bold(msg any, styles ...string) string {
	return bold(msg, styles, c)
}

func (c *Color) Dim(msg any, styles ...string) string {
	return dim(msg, styles, c)
}

func (c *Color) Italic(msg any, styles ...string) string {
	return italic(msg, styles, c)
}

func (c *Color) Underline(msg any, styles ...string) string {
	return underline(msg, styles, c)
}

func (c *Color) Inverse(msg any, styles ...string) string {
	return inverse(msg, styles, c)
}

func (c *Color) Hidden(msg any, styles ...string) string {
	return hidden(msg, styles, c)
}

func (c *Color) Strikeout(msg any, styles ...string) string {
	return strikeout(msg, styles, c)
}

// Output returns the output.
func Output() io.Writer {
	return global.output
}

// SetOutput sets the output.
func SetOutput(w io.Writer) {
	global.SetOutput(w)
}

func Disable() {
	global.Disable()
}

func Enable() {
	global.Enable()
}

// Print is analogous to `fmt.Print` with terminal detection.
func Print(args ...any) {
	global.Print(args...)
}

// Println is analogous to `fmt.Println` with terminal detection.
func Println(args ...any) {
	global.Println(args...)
}

// Printf is analogous to `fmt.Printf` with terminal detection.
func Printf(format string, args ...any) {
	global.Printf(format, args...)
}

func Black(msg any, styles ...string) string {
	return global.Black(msg, styles...)
}

func Red(msg any, styles ...string) string {
	return global.Red(msg, styles...)
}

func Green(msg any, styles ...string) string {
	return global.Green(msg, styles...)
}

func Yellow(msg any, styles ...string) string {
	return global.Yellow(msg, styles...)
}

func Blue(msg any, styles ...string) string {
	return global.Blue(msg, styles...)
}

func Magenta(msg any, styles ...string) string {
	return global.Magenta(msg, styles...)
}

func Cyan(msg any, styles ...string) string {
	return global.Cyan(msg, styles...)
}

func White(msg any, styles ...string) string {
	return global.White(msg, styles...)
}

func Gray(msg any, styles ...string) string {
	return global.Gray(msg, styles...)
}

func BlackBg(msg any, styles ...string) string {
	return global.BlackBg(msg, styles...)
}

func RedBg(msg any, styles ...string) string {
	return global.RedBg(msg, styles...)
}

func GreenBg(msg any, styles ...string) string {
	return global.GreenBg(msg, styles...)
}

func YellowBg(msg any, styles ...string) string {
	return global.YellowBg(msg, styles...)
}

func BlueBg(msg any, styles ...string) string {
	return global.BlueBg(msg, styles...)
}

func MagentaBg(msg any, styles ...string) string {
	return global.MagentaBg(msg, styles...)
}

func CyanBg(msg any, styles ...string) string {
	return global.CyanBg(msg, styles...)
}

func WhiteBg(msg any, styles ...string) string {
	return global.WhiteBg(msg, styles...)
}

func Reset(msg any, styles ...string) string {
	return global.Reset(msg, styles...)
}

func Bold(msg any, styles ...string) string {
	return global.Bold(msg, styles...)
}

func Dim(msg any, styles ...string) string {
	return global.Dim(msg, styles...)
}

func Italic(msg any, styles ...string) string {
	return global.Italic(msg, styles...)
}

func Underline(msg any, styles ...string) string {
	return global.Underline(msg, styles...)
}

func Inverse(msg any, styles ...string) string {
	return global.Inverse(msg, styles...)
}

func Hidden(msg any, styles ...string) string {
	return global.Hidden(msg, styles...)
}

func Strikeout(msg any, styles ...string) string {
	return global.Strikeout(msg, styles...)
}

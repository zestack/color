//go:build !windows

package color

import "github.com/mattn/go-isatty"

func Colorable(a any) bool {
	if f, ok := a.(interface{ Fd() uintptr }); ok {
		return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
	}
	return false
}

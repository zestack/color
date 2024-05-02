//go:build windows

package color

import (
	"github.com/mattn/go-isatty"
	"golang.org/x/sys/windows"
)

func enableVTMode(console windows.Handle) bool {
	mode := uint32(0)
	err := windows.GetConsoleMode(console, &mode)
	if err != nil {
		return false
	}

	// EnableVirtualTerminalProcessing is the console mode to allow ANSI code
	// interpretation on the console. See:
	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	// It only works on Windows 10. Earlier terminals will fail with an err which we will
	// handle to say don't color
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	err = windows.SetConsoleMode(console, mode)
	return err == nil
}

func Colorable(a any) bool {
	if f, ok := a.(interface{ Fd() uintptr }); ok {
		if isatty.IsTerminal(f.Fd()) {
			return enableVTMode(windows.Handle(f.Fd()))
		}
		return isatty.IsCygwinTerminal(f.Fd())
	}
	return false
}

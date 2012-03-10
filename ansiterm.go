// ansiterm.go

// see http://en.wikipedia.org/wiki/ANSI_escape_code
//		NOTE: won't work for older MSWin consoles - before DOS 2.0
//		NOTE: won't work for newer MSWin consoles - Win32 and later
//		Should work on your VT-100 if hardware still functions :-)
//		Works on (some/most?) xterms, including default GNOME consoles
//
//		Using "fmt" to send single bytes is probably not efficient
//		but unless you're sending massive data to screen you probably
//		won't notice...
//		It's primarily intended to make "status reports" easier to handle on
//		long running applications that don't warrant a full GUI workup.
//
//		See the demo program's Headline and StatusUpdate functions.
//		Because that's ALL it was intended to do I have implemented only
//		the ansi codes that were needed for that limited objective.

package ansiterm

import (
	"fmt"
)

const (
	ESC     = 033
	NORMAL  = 0
	INVERSE = 7
)

var (
	license = "Simplified BSD License, see README.md for details"
)

// erase whole page, leave cursor at 1,1
// 		ansi ED, special case n = 2
func ClearPage() {
	fmt.Printf("\033[2J")
	MoveToRC(1, 1)
}

// erase from cursor to end of line
// 		ansi EL specific case n = missing
func ClearLine() {
	fmt.Printf("\033[K")
}

// ansi SCP
func SavePosn() {
	fmt.Printf("\033[s")
}

// ansi RCP
func RestorePosn() {
	fmt.Printf("\033[u")
}

// erase N chars but dont move cursor position (clear field for printing)
func Erase(nchars int) {
	i := 0
	for nchars > 0 {
		nchars--
		fmt.Printf(" ")
		i++
	}
	for i > 0 {
		i--
		fmt.Printf("\b")
	}
}

// ansi HVP
func MoveToRC(row, col int) {
	fmt.Printf("\033[%d;%df", row, col)
}

// sugar for HVP
func MoveToXY(x, y int) {
	MoveToRC(y, x)
}

// 
func ResetTerm(attr int) {
	fmt.Printf("\033[1;80;0m") // restore normal attributes	
	if attr == NORMAL {
		return
	}
	fmt.Printf("033[1;1;%dm", attr)
}

package robot

import (
	"fmt"
	"testing"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
)

func TestGetCursorPos(t *testing.T) {

	EnumIndex = 0

	for EnumIndex < 1000 {
		var p api.POINT
		api.GetCursorPos(&p)
		p.X = 126
		p.Y = 104
		fmt.Println(p.X, p.Y)

		HWND := api.WindowFromPoint(p)
		HWND = 0x000D061E
		class := GetClassName(HWND)
		title := GetWindowText(HWND)
		// var p0 api.POINT
		// api.ClientToScreen(HWND, &p0)
		l, r, t, b := GetWindowRect(HWND)
		// c.Left += p0.X
		// c.Top += p0.Y
		fmt.Printf("%s,%s,0x%08X,%d,%d,%d,%d\n", title, class, HWND, l, t, r, b)
		robotgo.MicroSleep(1000)
		EnumIndex++
	}

}

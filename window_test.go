package robot

import (
	"fmt"
	"testing"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
	"github.com/ntcat/win32gui"
)

func TestGetCursorPos(t *testing.T) {

	EnumIndex = 0

	for EnumIndex < 1000 {
		var p api.POINT
		api.GetCursorPos(&p)
		p.X = 126
		p.Y = 104
		fmt.Println(p.X, p.Y)

		hWnd := api.WindowFromPoint(p)
		hWnd = 0x000D061E
		class := win32gui.GetClassName(hWnd)
		title := win32gui.GetWindowText(hWnd)
		// var p0 api.POINT
		// api.ClientToScreen(hWnd, &p0)
		l, r, t, b := win32gui.GetWindowRect(hWnd)
		// c.Left += p0.X
		// c.Top += p0.Y
		fmt.Printf("%s,%s,0x%08X,%d,%d,%d,%d\n", title, class, hWnd, l, t, r, b)
		robotgo.MicroSleep(1000)
		EnumIndex++
	}

}

// func TestGetCursorPos1(t *testing.T) {
// 	wl := NewTradeWin()
// 	wl.Init() //赋值HWnd、Opend
// 	EnumIndex = 0
// 	var p, p0 api.POINT
// 	api.ClientToScreen(wl.HWnd, &p0)
// 	handle := int(wl.HWnd)
// 	sHWnd := utils.DecimalToAny(handle, 16, 8)
// 	fmt.Println(sHWnd)
// 	if wl.Opend {
// 		for EnumIndex < 60 {
// 			api.GetCursorPos(&p)
// 			x := p.X //- p0.X
// 			y := p.Y // - p0.Y
// 			ctrl := wl.GetCtrlInWinPos(x, y)
// 			handle = int(ctrl.HWnd)
// 			sHWnd = utils.DecimalToAny(handle, 16, 8)
// 			fmt.Println(x, y)
// 			fmt.Println(ctrl.Title, ctrl.Class, sHWnd, ctrl.Rect)
// 			robotgo.MicroSleep(1000)
// 		}
// 	}

// }

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
		var c api.RECT
		api.GetWindowRect(HWND, &c)
		// c.Left += p0.X
		// c.Top += p0.Y
		fmt.Printf("%s,%s,0x%08X,%d,%d,%d,%d\n", title, class, HWND, c.Left, c.Top, c.Right, c.Bottom)
		robotgo.MicroSleep(1000)
		EnumIndex++
	}

}

// func TestGetCursorPos1(t *testing.T) {
// 	wl := NewTradeWin()
// 	wl.Init() //赋值HWND、Opend
// 	EnumIndex = 0
// 	var p, p0 api.POINT
// 	api.ClientToScreen(wl.HWND, &p0)
// 	handle := int(wl.HWND)
// 	sHWND := utils.DecimalToAny(handle, 16, 8)
// 	fmt.Println(sHWND)
// 	if wl.Opend {
// 		for EnumIndex < 60 {
// 			api.GetCursorPos(&p)
// 			x := p.X //- p0.X
// 			y := p.Y // - p0.Y
// 			ctrl := wl.GetCtrlInWinPos(x, y)
// 			handle = int(ctrl.HWND)
// 			sHWND = utils.DecimalToAny(handle, 16, 8)
// 			fmt.Println(x, y)
// 			fmt.Println(ctrl.Title, ctrl.Class, sHWND, ctrl.Rect)
// 			robotgo.MicroSleep(1000)
// 		}
// 	}

// }

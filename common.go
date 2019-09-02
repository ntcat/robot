package robot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
	"github.com/ntcat/win32gui"
)

// CloseWin Close Win
func CloseWin(hwnd api.HWND) {
	//PostMessage 发送关闭指令后立即退出。如果用SendMessage不关闭它不走
	api.PostMessage(hwnd, api.WM_CLOSE, 0, 0) //close windows
	robotgo.MicroSleep(200)
}

// PressButton press button
func PressButton(buttonHwnd api.HWND) {
	api.SendMessage(buttonHwnd, api.WM_LBUTTONDOWN, 0, 0) //鼠标左键按下
	api.SendMessage(buttonHwnd, api.WM_LBUTTONUP, 0, 0)   //鼠标左键抬起
}

// RectClick mouse click the center of rect
func RectClick(hWnd api.HWND) {
	x, y := win32gui.GetRectCenter(hWnd)
	if x > 0 {
		robotgo.MoveMouse(x, y)
		robotgo.Click()
	}
}

// ReadClip get string from clip
func ReadClip() string {
	text, err := robotgo.ReadAll()
	if err != nil {
		fmt.Println("robotgo.ReadAll err is: ", err)
		return ""
	}
	return text
}

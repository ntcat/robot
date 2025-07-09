package robot

import (
	"errors"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
)

// PressButton press button
func PressButton(buttonHWND api.HWND) {
	api.SendMessage(buttonHWND, api.WM_LBUTTONDOWN, 0, 0) //鼠标左键按下
	api.SendMessage(buttonHWND, api.WM_LBUTTONUP, 0, 0)   //鼠标左键抬起
}

// RectClick mouse click the center of rect
func RectClick(HWND api.HWND) {
	x, y := GetRectCenter(HWND)
	if x > 0 {
		robotgo.MoveMouse(x, y)
		robotgo.Click()
	}
}

// ReadClip get string from clip
func ReadClip() (string, error) {
	text, err := robotgo.ReadAll()
	if err != nil {
		return "", errors.New("robotgo.ReadAll err")
	}
	return text, nil
}

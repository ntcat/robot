package robot

import (
	"syscall"
	"unsafe"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
)

// EnumChildWindows 枚举所有子窗
func EnumChildWindows(hWndParent api.HWND, lpEnumFunc, lParam uintptr) bool {
	ret := api.EnumChildWindows(hWndParent, lpEnumFunc, lParam)
	return ret
}

// StringToUTF16PtrElseNil String To UTF16Ptr if empty string trans to nil
func StringToUTF16PtrElseNil(str string) *uint16 {
	if str == "" {
		return nil
	}
	return syscall.StringToUTF16Ptr(str)
}

// FindWindow find window hwnd by name class="" if nil,nil mean ignore it
func FindWindow(class, title string) api.HWND {
	var hwnd api.HWND
	hwnd = api.FindWindow(StringToUTF16PtrElseNil(class), StringToUTF16PtrElseNil(title))
	return hwnd
}

// FindWindowEx find window hwnd by name class="" if nil,nil mean ignore it
func FindWindowEx(hWndParent, hWndChild api.HWND, class, title string) api.HWND {
	var hwnd api.HWND
	hwnd = api.FindWindowEx(hWndParent, hWndChild,
		StringToUTF16PtrElseNil(class),
		StringToUTF16PtrElseNil(title))
	return hwnd
}

// GetClassName Get Class Name
func GetClassName(hWnd api.HWND) string {
	buf := make([]uint16, 255)
	ret, _ := api.GetClassName(hWnd, &buf[0], 255)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf)
}

// GetHWND get foreground window hwnd
func GetHWND() api.HWND {
	hwnd := api.GetForegroundWindow()

	return hwnd
}

// GetWindowText Get Window Text
func GetWindowText(hWnd api.HWND) string {
	textLength := api.SendMessage(hWnd, api.WM_GETTEXTLENGTH, 0, 0)
	buf := make([]uint16, textLength+1)
	api.SendMessage(hWnd, api.WM_GETTEXT, uintptr(textLength+1), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}

// SetWindowText Set Window Text
func SetWindowText(hWnd api.HWND, text string) {
	buf := syscall.StringToUTF16Ptr(text)
	api.SendMessage(hWnd, api.WM_SETTEXT, 0, uintptr(unsafe.Pointer(buf)))
}

// GetActiveWindow ...
func GetActiveWindow() api.HWND {
	return api.GetActiveWindow()
}

// SetActiveWindow ...
func SetActiveWindow(hWnd api.HWND) {
	api.SetActiveWindow(hWnd)
}

//BringWindowToTop ...
func BringWindowToTop(hWnd api.HWND) {
	api.BringWindowToTop(hWnd)
}

// CloseWin Close Win
func CloseWin(hwnd api.HWND) {
	//PostMessage 发送关闭指令后立即退出。如果用SendMessage不关闭它不走
	api.PostMessage(hwnd, api.WM_CLOSE, 0, 0) //close windows
	robotgo.MicroSleep(200)
}

// ShowWindow nCmdShow:
/*
api.SW_SHOWMAXIMIZED：
激活一个窗口并以最大化的状态显示它。

api.SW_SHOWMINIMIZED：
激活一个窗口并以最小化的状态显示它。

api.SW_SHOWMINNOACTIVE：
以最小化的状态来显示一个窗体，除非窗体是非激活状态的，否则本函数的效果类似于
SW_SHOWMINIMIZED。

api.SW_SHOWNA：
以当前的大小和位置来显示一个窗体，除非窗体是非激活状态的，否则本函数效果类似
于SW_SHOW。

api.SW_SHOWNOACTIVATE：
以最近的状态来显示一个窗体。除非窗台是非激活的，否则本函数的效果类似于
SW_SHOWNORMAL。

api.SW_SHOWNORMAL：
激活并显示一个窗体，如果此窗体是处于最大化或最小化的，将恢复为默认的大小和位置。
在程序第一次显示一个窗口时，应该设定这个标志。
*/
func ShowWindow(hwnd api.HWND, nCmdShow int32) {
	api.ShowWindow(hwnd, nCmdShow)
}

// GetWindowRect get Window Rect
func GetWindowRect(hWnd api.HWND) (int32, int32, int32, int32) {
	var rect api.RECT
	ok := api.GetWindowRect(hWnd, &rect)
	if ok {
		return rect.Left, rect.Right, rect.Top, rect.Bottom
	}
	return 0, 0, 0, 0
}

// GetClientRect get client rect
func GetClientRect(hWnd api.HWND) (int32, int32, int32, int32) {
	var rect api.RECT
	ok := api.GetClientRect(hWnd, &rect)
	if ok {
		return rect.Left, rect.Right, rect.Top, rect.Bottom
	}
	return 0, 0, 0, 0
}

// GetRectCenter get the center of rect
func GetRectCenter(hWnd api.HWND) (int, int) {
	var rect api.RECT
	ok := api.GetWindowRect(hWnd, &rect)
	if ok {
		l, r, t, b := rect.Left, rect.Right, rect.Top, rect.Bottom
		x := int(l + (r-l)/2)
		y := int(t + (b-t)/2)
		return x, y
	}
	return 0, 0
}

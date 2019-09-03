package robot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
	gui "github.com/ntcat/win32gui"
)

// EnumResoultCallBack 枚举结果回调函数，写成一个类型，有助于理解
type EnumResoultCallBack func(int, api.HWND, api.HWND, string, string, api.RECT)

var (
	// EnumIndex 全局变量，用于EnumWindow
	EnumIndex = 0
	// EnumResoult 枚举结果回调函数
	EnumResoult EnumResoultCallBack
)

// NewWindow 新建类
func NewWindow(title, class string) *Window {

	return &Window{
		Title:       title,
		Class:       class,
		HWnd:        0,
		Opend:       false,
		OpenTimeOut: 100,
		BeforeOpen:  func() {}, //can't set nil
		AfterOpen:   func() {},
		BeforeClose: func() {},
		AfterClose:  func() {},
	}

}

// Detect Detect HWnd
func (w *Window) Detect() *Window {
	w.HWnd = gui.FindWindow(w.Class, w.Title)
	if w.HWnd > 0 {
		w.Opend = true
	} else {
		w.Opend = false
		fmt.Printf("Can't detect window: %s\n", w.Title)
	}
	return w
}

// GetControls Get Controls of window, store in map
// 界面变动时，控件也会变化，就需要重新搜索重建一下
func (w *Window) GetControls() {
	if w.HWnd > 0 {
		w.Controls = make(map[int]Control) //放在EnumResoult中意味着每次调用都初始化
		EnumResoult = func(i int, hWnd, hWndParent api.HWND, class, title string, c api.RECT) {
			w.Controls[i] = Control{
				Index:      i,
				Title:      title,
				Class:      class,
				HWndParent: hWndParent,
				HWnd:       hWnd,
				Rect:       c,
			}
		}
		EnumWindow(w.HWnd)
	}
}

// GetControlByIndex GetControl by controls index
func (w *Window) GetControlByIndex(index int) Control {
	return w.Controls[index]
}

// GetControlByTitleClassRect GetControl by controls title, class & rect
// ltrb中如果要忽略某项，设置为负就可跳过.一般这二项就可以定位一个ctrl，如果不够，加个是否显示属性，那再说吧
func (w *Window) GetControlByTitleClassRect(title, class string, l, t, r, b int32) Control {
	var tBool, cBool, rBool bool

	for _, v := range w.Controls {
		if title == "" {
			tBool = true
		} else {
			tBool = (v.Title == title)
		}

		if class == "" {
			cBool = true
		} else {
			cBool = (v.Class == class)
		}

		if tBool && cBool {
			rBool = true
			if l > 0 {
				rBool = rBool && (v.Rect.Left == l)
			}
			if t > 0 {
				rBool = rBool && (v.Rect.Top == t)
			}
			if r > 0 {
				rBool = rBool && (v.Rect.Right == r)
			}
			if b > 0 {
				rBool = rBool && (v.Rect.Bottom == b)
			}
			if rBool {
				return v
			}

		}
	}
	return *NewControl()
}

// EnumWindow enum all window at any level
func EnumWindow(hWndParent api.HWND) {
	var child api.HWND
	var w, c api.RECT
	api.GetWindowRect(hWndParent, &w)
	for {
		child = api.FindWindowEx(hWndParent, child, nil, nil)

		//SetWindowText(child, strings.Repeat("12", index))
		if child == 0 {
			break
		} else {
			class := gui.GetClassName(child)
			title := gui.GetWindowText(child)
			api.GetWindowRect(child, &c)
			c.Left -= w.Left
			c.Top -= w.Top
			c.Right -= w.Left
			c.Bottom -= w.Top
			EnumResoult(EnumIndex, child, hWndParent, class, title, c)
			EnumIndex++
			EnumWindow(child)
		}

	}
}

// WaitOpen Wait Open
func (w *Window) WaitOpen() {
	var stop = false
	count := 0
	for !w.Opend && !stop {
		w.BeforeOpen()
		w.Detect()
		if count > w.OpenTimeOut {
			stop = true
		}
		robotgo.MicroSleep(200)
		count++
	}
	if !stop {
		w.AfterOpen()
	}
}

// PressButtonByTitle Press Button By Title
func (w *Window) PressButtonByTitle(title string) {
	for w.Opend {
		child := gui.FindWindowEx(w.HWnd, 0, "Button", title)
		if child > 0 {
			PressButton(child)
			w.Detect()
		} else {
			break
		}

	}
}

// Close Close tradeWin
func (w *Window) Close() {
	for w.Opend {
		w.BeforeClose()
		CloseWin(w.HWnd)
		w.AfterClose()
		w.Detect()
	}
}
func (w *Window) hasStyleBits(bits uint32) bool {
	return hasWindowLongBits(w.HWnd, api.GWL_STYLE, bits)
}

func (w *Window) hasExtendedStyleBits(bits uint32) bool {
	return hasWindowLongBits(w.HWnd, api.GWL_EXSTYLE, bits)
}

func hasWindowLongBits(hwnd api.HWND, index int32, bits uint32) bool {
	value := uint32(api.GetWindowLong(hwnd, index))

	return value&bits == bits
}

func setWindowVisible(hwnd api.HWND, visible bool) {
	var cmd int32
	if visible {
		cmd = api.SW_SHOWNA
	} else {
		cmd = api.SW_HIDE
	}
	api.ShowWindow(hwnd, cmd)
}

// BringToTop moves the *Window to the top of the keyboard focus order.
func (w *Window) BringToTop() uint32 {
	if !api.SetWindowPos(w.HWnd, api.HWND_TOP, 0, 0, 0, 0, api.SWP_NOACTIVATE|api.SWP_NOMOVE|api.SWP_NOSIZE) {
		return api.GetLastError()
	}

	return 0
}

// Focused returns whether the Window has the keyboard input focus.
func (w *Window) Focused() bool {
	return w.HWnd == api.GetFocus()
}

// SetFocus sets the keyboard input focus to the *Window.
func (w *Window) SetFocus() error {
	if api.SetFocus(w.HWnd) == 0 {
		return lastError("SetFocus")
	}

	return nil
}

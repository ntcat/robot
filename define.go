package robot

import api "github.com/ntcat/win32api"

// Control common Control struct
type Control struct {
	Index       int
	Description string
	Title       string
	Class       string
	HWndParent  api.HWND
	HWnd        api.HWND
	Rect        api.RECT
	/*是指在窗体中控件的索引号，不同于tabIndex,是指搜索枚举时的顺序，
	  可以是同类的顺序，也可以全部，看你如何使用它了，用来区分不同的控件
	*/
}

// Window common window struct
type Window struct {
	Title       string
	Class       string
	HWnd        api.HWND
	Rect        api.RECT
	Opend       bool
	Description string
	OpenTimeOut int
	visible     bool
	enabled     bool
	Controls    map[int]Control //它的值是动态变化的，界面不同呈现的控制也不同
	BeforeOpen  func()          // 窗口打开前的回调函数
	AfterOpen   func()          // 窗口打开后的回调函数
	BeforeClose func()          // 窗口关闭前的回调函数
	AfterClose  func()          // 窗口关闭后的回调函数
}

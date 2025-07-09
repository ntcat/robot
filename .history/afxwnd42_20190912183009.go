package robot

import (
	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
)

// AfxWnd42Ctrl AfxWnd42 Ctrl
type AfxWnd42Ctrl struct {
	Control
}

// NewAfxWnd42  新建类
func NewAfxWnd42(ctrl Control) *AfxWnd42Ctrl {
	return &AfxWnd42Ctrl{
		Control: ctrl,
	}
}

// Init init object
func (c *AfxWnd42Ctrl) Init(description string) *AfxWnd42Ctrl {
	if c != nil {
		c.Description = description
		return c
	}
	return nil
}

// SetValue set textbox value
func (c *AfxWnd42Ctrl) SetValue(v string) error {
	//SetWindowText(HWND, code) //这句不起作用,AfxWnd42特殊的很
	if c != nil && c.HWND > 0 {
		api.SetFocus(c.HWND)
		robotgo.WriteAll(v)         //复制到剪切板
		robotgo.KeyTap("v", "ctrl") //粘贴
		return nil
	}
	return errHWND
}

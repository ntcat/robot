package robot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
)

// AfxWnd42Ctrl AfxWnd42 Ctrl
type AfxWnd42Ctrl struct {
	Control
}

// NewAfxWnd42Ctrl 新建类
func NewAfxWnd42Ctrl(ctrl Control) *AfxWnd42Ctrl {
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
	fmt.Printf("error info - %s hwnd：%08X\n", c.Description, c.HWnd)
	return nil
}

// SetValue set textbox value
func (c *AfxWnd42Ctrl) SetValue(v string) {
	//SetWindowText(hWnd, code) //这句不起作用,AfxWnd42特殊的很
	if c != nil && c.HWnd > 0 {
		api.SetFocus(c.HWnd)
		robotgo.WriteAll(v)         //复制到剪切板
		robotgo.KeyTap("v", "ctrl") //粘贴
	} else {
		fmt.Printf("Can't SetValue - %s hwnd：%08X\n", c.Description, c.HWnd)
	}

}

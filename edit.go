package robot

import (
	"fmt"

	"github.com/ntcat/win32gui"
)

// EditCtrl Edit Ctrl
type EditCtrl struct {
	Control
}

// NewEditCtrl 新建类
func NewEditCtrl(ctrl Control) *EditCtrl {
	return &EditCtrl{
		Control: ctrl,
	}
}

// Init init object
func (c *EditCtrl) Init(description string) *EditCtrl {
	if c != nil {
		c.Description = description
		return c
	}
	fmt.Printf("error info - %s hwnd：%08X\n", c.Description, c.HWnd)
	return nil
}

// SetValue set textbox value
func (c *EditCtrl) SetValue(v string) {
	if c != nil && c.HWnd > 0 {
		win32gui.SetWindowText(c.HWnd, v)
	} else {
		fmt.Printf("Can't SetValue - %s hwnd：%08X\n", c.Description, c.HWnd)
	}
}

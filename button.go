package robot

import (
	"fmt"
)

// ButtonCtrl Button Ctrl
type ButtonCtrl struct {
	Control
}

// NewButtonCtrl 新建类
func NewButtonCtrl(ctrl Control) *ButtonCtrl {
	return &ButtonCtrl{
		Control: ctrl,
	}
}

// Init init object
func (c *ButtonCtrl) Init(description string) *ButtonCtrl {
	if c != nil {
		c.Description = description
		return c
	}
	fmt.Printf("error info - %s hwnd：%08X\n", c.Description, c.HWnd)
	return nil
}

// Click Click
func (c *ButtonCtrl) Click() {
	if c != nil && c.HWnd > 0 {
		RectClick(c.HWnd)
	}
}

// ClickUntilEnd Click Until something end
func (c *ButtonCtrl) ClickUntilEnd() {
	if c != nil && c.HWnd > 0 {
		PressButton(c.HWnd)
	}
}

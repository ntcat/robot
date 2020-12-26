package robot

// ButtonCtrl Button Ctrl
type ButtonCtrl struct {
	Control
}

// NewButton 新建类
func NewButton(ctrl Control) *ButtonCtrl {
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
	return nil
}

// Click Click
func (c *ButtonCtrl) Click() error {
	if c != nil && c.HWND > 0 {
		RectClick(c.HWND)
		return nil
	}
	return ErrHWND
}

// ClickUntilEnd Click Until something end
func (c *ButtonCtrl) ClickUntilEnd() error {
	if c != nil && c.HWND > 0 {
		PressButton(c.HWND)
		return nil
	}
	return ErrHWND
}

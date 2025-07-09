package robot

// EditCtrl Edit Ctrl
type EditCtrl struct {
	Control
}

// NewEdit 新建类
func NewEdit(ctrl Control) *EditCtrl {
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
	return nil
}

// SetValue set textbox value
func (c *EditCtrl) SetValue(v string) error {
	if c != nil && c.HWND > 0 {
		SetWindowText(c.HWND, v)
		return nil
	}
	return errHWND
}

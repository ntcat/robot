package robot

import (
	"fmt"

	api "github.com/ntcat/win32api"
)

// ListViewCtrl ListView Ctrl
type ListViewCtrl struct {
	Control
	SelectIndex string
	ItemHight   int
}

// NewListView 新建类
func NewListView(ctrl Control) *ListViewCtrl {
	return &ListViewCtrl{
		Control:   ctrl,
		ItemHight: 20,
	}
}

// Init init object
func (lv *ListViewCtrl) Init(description string) *ListViewCtrl {
	if lv != nil {
		lv.Description = description
		return lv
	}
	fmt.Printf("error info - %s hwnd：%08X\n", lv.Description, lv.HWnd)
	return nil
}

// GetRows get listView items count
func (lv *ListViewCtrl) GetRows() int {
	if lv.HWnd == 0 {
		fmt.Printf("error info - %s hwnd：%08X\n", lv.Description, lv.HWnd)
		return 0
	}
	rows := api.SendMessage(lv.HWnd, api.LVM_GETITEMCOUNT, 0, 0)
	return int(rows)
}

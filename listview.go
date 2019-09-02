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

// GetCols get listView items cloumns
func (lv *ListViewCtrl) GetCols() int {
	if lv.HWnd == 0 {
		fmt.Printf("error info - %s hwnd：%08X\n", lv.Description, lv.HWnd)
		return 0
	}
	//第二个参数是 LVM_GETHEADER,获得LISTVIEW的HEADER句柄
	var lngHeaderHwnd api.HWND
	lngHeaderHwnd = (api.HWND)(api.SendMessage(lv.HWnd, api.LVM_GETHEADER, 0, 0))
	if lngHeaderHwnd > 0 {
		lngCols := api.SendMessage(lngHeaderHwnd, api.HDM_GETITEMCOUNT, 0, 0) //获取ListView表头项目数
		return int(lngCols)
	}
	return 1
}

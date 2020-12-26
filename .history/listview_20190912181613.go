package robot

import (
	"errors"

	api "github.com/ntcat/win32api"
)

// ListViewCtrl ListViewCtrl Ctrl
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
	return nil
}

// GetRows get ListViewCtrl items count
func (lv *ListViewCtrl) GetRows() int {
	if lv.HWND == 0 {
		return 0
	}
	rows := api.SendMessage(lv.HWND, api.LVM_GETITEMCOUNT, 0, 0)
	return int(rows)
}

// GetCols get ListViewCtrl items cloumns
func (lv *ListViewCtrl) GetCols() (int, error) {
	if lv.HWND == 0 {
		return 0, errHWND
	}
	//第二个参数是 LVM_GETHEADER,获得ListViewCtrl的HEADER句柄
	var lngHeaderHWND api.HWND
	lngHeaderHWND = (api.HWND)(api.SendMessage(lv.HWND, api.LVM_GETHEADER, 0, 0))
	if lngHeaderHWND > 0 {
		lngCols := api.SendMessage(lngHeaderHWND, api.HDM_GETITEMCOUNT, 0, 0) //获取ListViewCtrl表头项目数
		return int(lngCols), nil
	}
	return 0, errors.New("lngHeaderHWND=0")
}

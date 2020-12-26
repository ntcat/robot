package robot

import (
	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
)

// TreeViewCtrl TreeView Ctrl
type TreeViewCtrl struct {
	Control
	Items     map[string][]int
	ItemHight int
	IconWidth int
}

// NewTreeView 新建类
func NewTreeView(ctrl Control) *TreeViewCtrl {
	return &TreeViewCtrl{
		Control:   ctrl,
		ItemHight: 20,
		IconWidth: 30,
		Items:     map[string][]int{}, //没有这个初始化，AddItem be error

	}
}

// GetItemHeight get Item Height,=20
func (tv *TreeViewCtrl) GetItemHeight() int {
	tv.ItemHight = int(api.SendMessage(tv.HWND, api.TVM_GETITEMHEIGHT, 0, 0))
	return tv.ItemHight
}

// Init init object
func (tv *TreeViewCtrl) Init(description string) *TreeViewCtrl {
	if tv != nil {
		tv.Description = description
		tv.ItemHight = tv.GetItemHeight()
		return tv
	}
	return nil
}

// AddItem 格式：AddItem("当日成交",[]int{9,3}) 代表第一级菜单在第9个，第二级在第3个
func (tv *TreeViewCtrl) AddItem(key string, value []int) {
	tv.Items[key] = value
}

// ClickItem 点击TV项
func (tv *TreeViewCtrl) ClickItem(key string) error {
	if tv == nil || tv.HWND == 0 {
		return errHWND
	}

	var p api.POINT
	api.ClientToScreen(tv.HWND, &p)
	ItemIndexs := tv.Items[key]
	//展开所有级别，同时点击
	var allV = 0
	for i, v := range ItemIndexs {
		allV += v
		x := int(p.X+tv.Rect.Left) + tv.IconWidth*(i+1)
		y := int(p.Y+tv.Rect.Top) + int(float64(tv.ItemHight)*(float64(allV)-0.5))
		robotgo.MoveMouse(x, y)
		robotgo.Click("left", false) //单击鼠标左键
		robotgo.MicroSleep(500)
	}
	robotgo.MicroSleep(1000)
	//单击折叠第一级，防止影响以后的操作
	x := int(p.X+tv.Rect.Left) + tv.IconWidth*(0+1)
	y := int(p.Y+tv.Rect.Top) + int(float64(tv.ItemHight)*(float64(ItemIndexs[0])-0.5))
	robotgo.MoveMouse(x, y)
	robotgo.Click("left", false) //单击鼠标左键
	robotgo.MicroSleep(500)
	return nil
}

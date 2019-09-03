package robot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	api "github.com/ntcat/win32api"
)

/*
这里是treeView的控制和操作代码
要求：
1、要点击的treeView项，必须出现在屏幕上，不能超出，也不能被其它窗口盖住，
1.5、也不要期望软件会帮你拖动滚动条
2、这里默认每个项的高度是20，好象改变屏幕分辨率也不影响
3、第一级左边距是30，第二级是60，看上去没问题

*/

// TreeView TreeView Ctrl
type TreeView struct {
	Control
	Items     map[string][]int
	ItemHight int
	IconWidth int
}

// NewTreeView 新建类
func NewTreeView(ctrl Control) *TreeView {
	return &TreeView{
		Control:   ctrl,
		ItemHight: 20,
		IconWidth: 30,
		Items:     map[string][]int{}, //没有这个初始化，AddItem be error

	}
}

// GetItemHeight get Item Height,=20
func (tv *TreeView) GetItemHeight() int {
	tv.ItemHight = int(api.SendMessage(tv.HWnd, api.TVM_GETITEMHEIGHT, 0, 0))
	return tv.ItemHight
}

// Init init object
func (tv *TreeView) Init() *TreeView {
	if tv != nil {
		tv.AddItem("买入", []int{0})
		tv.AddItem("卖出", []int{1})
		tv.AddItem("市价买入", []int{2})
		tv.AddItem("市价卖出", []int{3})
		tv.AddItem("资金股份", []int{8, 0})
		tv.AddItem("当日成交", []int{8, 3})
		tv.AddItem("交割单", []int{8, 8})
		return tv
	}
	fmt.Printf("error info - %s hwnd：%08X\n", tv.Description, tv.HWnd)
	return nil
}

// AddItem 格式：AddItem("当日成交",[]int{9,3}) 代表第一级菜单在第9个，第二级在第3个
func (tv *TreeView) AddItem(key string, value []int) {
	tv.Items[key] = value
}

// ClickItem 点击TV项
func (tv *TreeView) ClickItem(key string) {
	if tv == nil || tv.HWnd == 0 {
		fmt.Printf("error TreeView ClickItem - %s hwnd：%08X\n", tv.Description, tv.HWnd)
		return
	}

	var p api.POINT
	api.ClientToScreen(tv.HWnd, &p)
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
}

package robot

import (
	"fmt"
	"testing"
)

func TestTreeView(t *testing.T) {
	//打开记事本 - 文件 -  打印
	w := NewWindow("通达信网上交易V6 四川分公司 申志红", "")
	w.Detect().GetControls()
	ctrl := w.GetControlByTitleClassRect("", "SysTreeView32", 2, 24, -1, -1)
	tv := NewTreeView(ctrl)
	h := tv.ItemHeight()
	fmt.Printf("hight:%d\n", h)
}
func TestTreeViewClick(t *testing.T) {
	//打开记事本 - 文件 -  打印
	w := NewWindow("通达信网上交易V6 四川分公司 申志红", "")
	w.Detect().GetControls()
	ctrl := w.GetControlByTitleClassRect("", "SysTreeView32", 2, 24, -1, -1)
	tv := NewTreeView(ctrl)
	var processID uint32
	api.GetWindowThreadProcessId(tv.HWnd, &processID)
	process := OpenProcess(PROCESS_ALL_ACCESS, FALSE, PID)
	fmt.Printf("hight:%d\n", h)
}

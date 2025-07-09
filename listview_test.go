package robot

import (
	"fmt"
	"testing"
)

func TestListView(t *testing.T) {
	//打开记事本 - 文件 -  打印
	w := NewWindow("打印", "")
	w.Detect().GetControls()
	ctrl := w.GetControlByTitleClassRect("FolderView", "SysListView32", -1, -1, -1, -1)
	lv := NewListView(ctrl)

	lv = lv.Init("test")
	rows := lv.GetRows()
	cols, _ := lv.GetCols()
	fmt.Printf("Row:%d - Col:%d\n", rows, cols)
}

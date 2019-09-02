package robot

import (
	"fmt"
	"testing"
)

func TestListView(t *testing.T) {
	//StartApp()
	w := NewWindow("block_select_dlg", "")
	w.Detect().GetControls()
	ctrl := w.GetControlByTitleClassRect("list2", "", -1, -1, -1, -1)
	lv := NewListView(ctrl)
	rows := lv.Init("test").GetRows()
	fmt.Printf("%d", rows)
}

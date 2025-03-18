package keyboardlayout

import (
	"fmt"
	"testing"
)

func TestKeyLayout(t *testing.T) {
	id := GetForegroundWindowKeyboardLayout()
	fmt.Printf("GetForegroundWindowKeyboardLayout: %x\n", id)
	l := GetKeyboardLayoutList()
	str := GetKeyboardLayoutName()
	fmt.Printf("%s\n", str)
	//ChangeForegroundWindowKeyboardLayout(l[0])
	str = GetKeyboardLayoutName()
	fmt.Printf("%s\n", str)
	//ChangeForegroundWindowKeyboardLayout(0)
	str = GetKeyboardLayoutName()
	fmt.Printf("%s\n", str)
	n := GetForegroundWindowKeyboardLayout()
	fmt.Printf("%d\n", n)
	var stringLayoutID = "4090409"
	var flags uint32 = 0x109
	m := LoadKeyboardLayout(stringLayoutID, flags)
	fmt.Printf("LoadKeyboardLayout: %x\n", id)
	fmt.Printf("0x%x-0x%x-0x%x-0x%x-0x%x-0x%x-0x%x\n%X\n", l[0], l[1], l[2], l[3], l[4], l[5], l[6], m)
}

func TestChangeLan(t *testing.T) {
	// timeUnix := time.Now().Unix() //已知的时间戳
	// formatTimeStr1 := time.Now().Format("2006-01-02") + " 09:30:00"
	// formatTimeStr := time.Unix(timeUnix, 8).Format("2006-01-02 15:04:05")
	// fmt.Println(formatTimeStr1) //打
	// fmt.Println(formatTimeStr)  //打

	// date := time.Now().Format("2006-01-02")
	// Today0930, _ := time.Parse("2006-01-02 15:04:05", date+" 09:30:00")
	// if Today0930.Before(Today0930) {
	// 	fmt.Println(Today0930) //打
	// }
	formatTimeStr := 0 - 0%3
	fmt.Println(formatTimeStr)

}

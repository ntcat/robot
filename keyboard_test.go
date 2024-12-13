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
	//ChangeLan(0x8040804) //ZH
	ChangeLan(0x4090409) //EN
	//ChangeLan(0x04190419) //russian
}

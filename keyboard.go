package keyboardlayout

import (
	"syscall"

	api "gorobot/win32api"
)

// GetForegroundWindowKeyboardLayout Get Foreground Window Keyboard Layout
func GetForegroundWindowKeyboardLayout() uint32 {
	hWnd := api.GetForegroundWindow()
	var threadID uint32
	api.GetWindowThreadProcessId(hWnd, &threadID)
	layoutID := api.GetKeyboardLayout(threadID)
	return layoutID
}

/*
Change foreground window keyboard layout
Parameter
	layout_id=0 : integer, 0 is english
	Integer containing a locale id, eg 68748313 - 0x04190419 - 0x419 - russian
	By default change layout like Ctrl+Shift or Alt+Shift
Return Value
Returns True if layout is changed and win32api.SendMessage() output if not
*/
// ChangeForegroundWindowKeyboardLayout
func ChangeForegroundWindowKeyboardLayout(layoutID uint32) bool {
	hWnd := api.GetForegroundWindow()
	result := api.SendMessage(hWnd, api.WM_INPUTLANGCHANGEREQUEST, 0, uintptr(layoutID))
	if result == 0 {
		return true
	}
	return false
}

// GetKeyboardLayoutList return array like [343434,54656]
func GetKeyboardLayoutList() []uint32 {
	const MaxLays = 10
	lpList := make([]uint32, MaxLays)
	n := api.GetKeyboardLayoutList(0, nil)
	n = api.GetKeyboardLayoutList(n, &lpList[0])
	return lpList
}

// GetKeyboardLayoutName Get Keyboard Layout Name
func GetKeyboardLayoutName() string {
	buf := make([]uint16, 255)
	ret := api.GetKeyboardLayoutName(&buf[0])
	if ret == false {
		return ""
	}
	return syscall.UTF16ToString(buf)
}

// LoadKeyboardLayout Load Keyboard Layout
func LoadKeyboardLayout(stringLayoutID string, flags uint32) uint32 {
	ret := api.LoadKeyboardLayout(syscall.StringToUTF16Ptr(stringLayoutID), flags)
	return uint32(ret)
}

// ChangeLan change language
/*
   EN = 0x4090409
   ZH = 0x8040804
*/
func ChangeLan(layoutID uint32) {

	klList := GetKeyboardLayoutList()

	var IsKeyInRange = func(key uint32, list []uint32) bool {
		for _, id := range list {
			if id == key {
				return true
			}
		}
		return false
	}
	if !IsKeyInRange(layoutID, klList) {
		LoadKeyboardLayout("4090409", 0)
	}
	ChangeForegroundWindowKeyboardLayout(layoutID)
}

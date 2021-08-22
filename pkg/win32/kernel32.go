package win32

import (
	"syscall"
)

var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultLCID = kernel32.NewProc("GetUserDefaultLCID")
	procGetConsoleWindow   = kernel32.NewProc("GetConsoleWindow")
)

func GetUserDefaultLCID() uint32 {
	ret, _, _ := procGetUserDefaultLCID.Call()
	return uint32(ret)
}

func GetConsoleWindow() HWND {
	ret, _, _ := procGetConsoleWindow.Call()
	return HWND(ret)
}

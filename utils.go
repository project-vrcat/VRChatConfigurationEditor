package main

import (
	"syscall"

	"github.com/TheTitanrain/w32"
)

var (
	user32          = syscall.MustLoadDLL("user32.dll")
	procEnumWindows = user32.MustFindProc("EnumWindows")
)

// IsChineseSimplified 当前操作系统语言是否为简体中文
func IsChineseSimplified() bool {
	switch w32.GetUserDefaultLCID() {
	case 0x0804:
		return true
	}
	return false
}

// FindWindowByProcessId 通过PID寻找窗口
func FindWindowByProcessId(pid int) w32.HWND {
	_hwnd := w32.HWND(0)
	cb := syscall.NewCallback(func(hwnd syscall.Handle, lParam uintptr) uintptr {
		_, _pid := w32.GetWindowThreadProcessId(w32.HWND(hwnd))
		if _pid == pid && w32.IsWindowVisible(w32.HWND(hwnd)) {
			_hwnd = w32.HWND(hwnd)
			return 0
		}
		return 1
	})
	_ = EnumWindows(cb, 0)
	return _hwnd
}

// EnumWindows 枚举窗口
func EnumWindows(enumFunc uintptr, lParam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, enumFunc, lParam, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

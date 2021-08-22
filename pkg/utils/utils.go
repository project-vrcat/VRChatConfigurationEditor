package utils

import (
	"os"

	"github.com/project-vrcat/VRChatConfigurationEditor/pkg/win32"
)

// IsChineseSimplified 当前操作系统语言是否为简体中文
func IsChineseSimplified() bool {
	switch win32.GetUserDefaultLCID() {
	case 0x0804:
		return true
	}
	return false
}

// HideConsoleWindow 隐藏控制台窗口
func HideConsoleWindow() {
	hwnd := win32.GetConsoleWindow()
	if hwnd <= 0 {
		return
	}
	_, pid := win32.GetWindowThreadProcessId(hwnd)
	if pid == os.Getpid() {
		win32.ShowWindow(hwnd, 0)
	}
}

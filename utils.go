package main

import (
	"github.com/TheTitanrain/w32"
)

// IsChineseSimplified 当前操作系统语言是否为简体中文
func IsChineseSimplified() bool {
	switch w32.GetUserDefaultLCID() {
	case 0x0804:
		return true
	}
	return false
}

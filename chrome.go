package main

import (
	"os/exec"

	"github.com/TheTitanrain/w32"
)

// PromptDownload 弹出询问是否下载Chrome对话框
func PromptDownload() {
	title := "Chrome not found"
	message := "No Chrome/Chromium installation was found. Would you like to download and install it now?"
	downloadUrl := "https://www.google.com/chrome/"
	downloadUrlChina := "https://www.google.cn/chrome/"

	r := w32.MessageBox(w32.HWND(0), message, title, w32.MB_YESNO|w32.MB_ICONQUESTION)
	if r == w32.IDYES {
		switch w32.GetUserDefaultLCID() {
		case 0x0804: // 如果系统语言为简体中文, 使用中国专用Chrome下载链接
			downloadUrl = downloadUrlChina
		}
		_ = exec.Command("cmd", "/c", "start", downloadUrl).Start()
	}
}

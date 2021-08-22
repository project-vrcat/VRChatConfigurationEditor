package app

import (
	"os/exec"

	"github.com/project-vrcat/VRChatConfigurationEditor/pkg/utils"
	"github.com/project-vrcat/VRChatConfigurationEditor/pkg/win32"
)

// PromptDownload 弹出询问是否下载Chrome对话框
func PromptDownload() {
	title := "Chrome not found"
	message := "No Chrome/Chromium installation was found. Would you like to download and install it now?"
	downloadUrl := "https://www.google.com/chrome/"
	downloadUrlChina := "https://www.google.cn/chrome/"

	r := win32.MessageBox(0, message, title, win32.MB_YESNO|win32.MB_ICONQUESTION)
	if r == win32.IDYES {
		// 如果系统语言为简体中文, 使用中国专用Chrome下载链接
		if utils.IsChineseSimplified() {
			downloadUrl = downloadUrlChina
		}
		_ = exec.Command("cmd", "/c", "start", downloadUrl).Start()
	}
}

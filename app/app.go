package app

import (
	"fmt"

	"github.com/project-vrcat/VRChatConfigurationEditor/pkg/utils"
	"github.com/project-vrcat/VRChatConfigurationEditor/pkg/win32"
	"github.com/zserge/lorca"
)

var ui lorca.UI

func init() {
	utils.HideConsoleWindow()
}

func bind() {
	_ = ui.Bind("vrchatPath", BindVRChatPath)
	_ = ui.Bind("readTextFile", BindReadTextFile)
	_ = ui.Bind("writeTextFile", BindWriteTextFile)
	_ = ui.Bind("selectDirectory", BindSelectDirectory)
	_ = ui.Bind("removeAll", BindRemoveAll)
	_ = ui.Bind("appVersion", BindAppVersion)
	_ = ui.Bind("open", BindOpen)
	_ = ui.Bind("checkUpdate", BindCheckUpdate)
}

// Run 运行App
func Run() {
	port := server()
	if lorca.ChromeExecutable() == "" {
		PromptDownload()
		return
	}
	var err error
	ui, err = lorca.New("", "", 800, 500)
	if err != nil {
		_ = win32.MessageBox(win32.HWND(0), err.Error(), "Error", win32.MB_OK|win32.MB_ICONERROR)
		return
	}
	defer ui.Close()

	bind()
	_ = ui.Load(fmt.Sprintf("http://127.0.0.1:%d", port))
	<-ui.Done()
}

//go:generate go-winres simply --icon public/favicon.png
package main

import (
	"fmt"

	"github.com/TheTitanrain/w32"
	"github.com/zserge/lorca"
)

func main() {
	port := server()
	if lorca.ChromeExecutable() == "" {
		PromptDownload()
		return
	}
	ui, err := lorca.New("", "", 800, 500)
	if err != nil {
		w32.MessageBox(w32.HWND(0), err.Error(), "Error", w32.MB_OK|w32.MB_ICONERROR)
		return
	}
	defer ui.Close()

	_ = ui.Bind("setWindowTitle", BindSetWindowTitle)
	_ = ui.Bind("vrchatPath", BindVRChatPath)
	_ = ui.Bind("readTextFile", BindReadTextFile)
	_ = ui.Bind("writeTextFile", BindWriteTextFile)
	_ = ui.Bind("selectDirectory", BindSelectDirectory)
	_ = ui.Bind("removeAll", BindRemoveAll)
	_ = ui.Bind("appVersion", BindAppVersion)
	_ = ui.Bind("open", BindOpen)
	_ = ui.Bind("checkUpdate", BindCheckUpdate)

	_ = ui.Load(fmt.Sprintf("http://127.0.0.1:%d", port))

	<-ui.Done()
}

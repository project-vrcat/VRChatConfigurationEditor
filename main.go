//go:generate go-winres simply --icon public/favicon.png
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/url"

	"github.com/TheTitanrain/w32"
	"github.com/zserge/lorca"
)

func main() {
	port := server()
	if lorca.ChromeExecutable() == "" {
		PromptDownload()
		return
	}
	loadingFile, _ := publicFiles.ReadFile("public/loading.html")
	ui, err := lorca.New("data:text/html,"+url.PathEscape(string(loadingFile)), "", 800, 500)
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

	_ = ui.Load(fmt.Sprintf("http://127.0.0.1:%d", port))

	<-ui.Done()
}

func server() int {
	port := pickPort()
	http.Handle("/", http.StripPrefix("/", http.FileServer(getFileSystem())))
	go func() {
		_ = http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil)
	}()
	return port
}

func pickPort() int {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		return -1
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

//go:embed public
var publicFiles embed.FS

func getFileSystem() http.FileSystem {
	fSys, _ := fs.Sub(publicFiles, "public")
	return http.FS(fSys)
}

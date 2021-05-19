//go:generate go-winres simply --icon public/favicon.png
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/url"

	"github.com/sqweek/dialog"
	"github.com/zserge/lorca"
)

func main() {
	port := server()
	if lorca.LocateChrome() == "" {
		ErrorMessage("Chrome installation is not found.")
		return
	}
	loadingFile, _ := publicFiles.ReadFile("public/loading.html")
	ui, err := lorca.New("data:text/html,"+url.PathEscape(string(loadingFile)), "", 800, 500)
	if err != nil {
		ErrorMessage(err.Error())
		return
	}
	defer ui.Close()

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

func ErrorMessage(msg string) {
	dialog.
		Message("%s", msg).
		Title("Error").
		Error()
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

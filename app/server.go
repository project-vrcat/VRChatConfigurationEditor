package app

import (
	"embed"
	"fmt"
	"io/fs"
	"mime"
	"net"
	"net/http"
	"os"
)

// PublicFiles 嵌入资源
var PublicFiles embed.FS

func init() {
	// Good job Microsoft :)
	// https://github.com/golang/go/issues/32350
	_ = mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
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

func getFileSystem() http.FileSystem {
	if len(os.Args) > 1 && (os.Args[1] == "live" || os.Args[1] == "--live") {
		return http.Dir("./public")
	}
	fSys, _ := fs.Sub(PublicFiles, "public")
	return http.FS(fSys)
}

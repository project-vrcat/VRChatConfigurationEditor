package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
)

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

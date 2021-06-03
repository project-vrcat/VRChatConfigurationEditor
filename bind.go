package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/TheTitanrain/w32"
)

// BindVRChatPath 获取VRChat配置目录
func BindVRChatPath() (_path string, err error) {
	appdata := os.Getenv("AppData")
	if appdata == "" {
		err = errors.New("The AppData environment variable must be set for app to run correctly.")
		return
	}
	if strings.Contains(appdata, "\\Roaming") {
		appdata = filepath.Dir(appdata)
	}
	_path = filepath.Join(appdata, "LocalLow\\VRChat\\VRChat")
	_, err = os.Stat(_path)
	return
}

// BindReadTextFile 读取文本格式的文件
func BindReadTextFile(filename string) (content string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	content = string(data)
	return
}

// BindWriteTextFile 写入文本格式的文件
func BindWriteTextFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)
}

// BindSelectDirectory 弹出目录选择框
func BindSelectDirectory(title string) (string, error) {
	pid := ui.Getpid()
	owner := FindWindowByProcessId(pid)
	_title, _ := syscall.UTF16PtrFromString(title)

	res := w32.SHBrowseForFolder(&w32.BROWSEINFO{
		Owner: owner,
		Flags: w32.BIF_RETURNONLYFSDIRS | w32.BIF_NEWDIALOGSTYLE,
		Title: _title,
	})
	if res == 0 {
		return "", errors.New("Cancelled")
	}
	return w32.SHGetPathFromIDList(res), nil
}

// BindRemoveAll 清空指定目录
func BindRemoveAll(path string) error {
	return os.RemoveAll(path)
}

// BindAppVersion 获取应用版本号及编译日期
func BindAppVersion() string {
	return fmt.Sprintf("%s build-%s", Version, BuildDate)
}

// BindOpen 通过系统默认浏览器打开指定url
func BindOpen(url string) error {
	return exec.Command("cmd", "/c", "start", url).Start()
}

// BindCheckUpdate 检查是否存在新版本
func BindCheckUpdate() bool {
	api := "https://api.github.com/repos/project-vrcat/VRChatConfigurationEditor/releases/latest"
	// GitHub API的IP https://api.github.com/meta
	// 防止因DNS污染而无法获取信息
	githubApiIp := "18.179.245.253"
	client := http.Client{Timeout: time.Second * 10}
	if IsChineseSimplified() {
		client.Transport = &http.Transport{
			DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
				_, port, err := net.SplitHostPort(addr)
				if err != nil {
					return nil, err
				}
				var dialer net.Dialer
				return dialer.DialContext(ctx, network, net.JoinHostPort(githubApiIp, port))
			},
		}
	}
	req, _ := http.NewRequest(http.MethodGet, api, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	var info struct {
		TagName string `json:"tag_name"`
	}
	if json.Unmarshal(data, &info) != nil {
		return false
	}
	return info.TagName != Version
}

package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/project-vrcat/VRChatConfigurationEditor/pkg/win32"
)

// BindVRChatPath 获取VRChat配置目录
func BindVRChatPath() (_path string, err error) {
	var localLow string
	appdata := os.Getenv("LOCALAPPDATA")
	if appdata == "" {
		appdata = os.Getenv("APPDATA")
	}
	if appdata == "" {
		err = errors.New("the AppData environment variable must be set for app to run correctly")
		return
	}
	localLow = filepath.Join(appdata, "../LocalLow")
	_path = filepath.Join(localLow, "VRChat/VRChat")
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
	owner := win32.FindWindowByProcessId(pid)
	_title, _ := syscall.UTF16PtrFromString(title)

	res := win32.SHBrowseForFolder(&win32.BROWSEINFO{
		Owner: owner,
		Flags: win32.BIF_RETURNONLYFSDIRS | win32.BIF_NEWDIALOGSTYLE,
		Title: _title,
	})
	if res == 0 {
		return "", errors.New("cancelled")
	}
	return win32.SHGetPathFromIDList(res), nil
}

// BindRemoveAll 清空指定目录
func BindRemoveAll(path string) error {
	return os.RemoveAll(path)
}

// BindAppVersion 获取应用版本号及编译日期
func BindAppVersion() string {
	return fmt.Sprintf("%s build-%s", Version, GitHash)
}

// BindOpen 通过系统默认浏览器打开指定url
func BindOpen(url string) error {
	return exec.Command("cmd", "/c", "start", url).Start()
}

// BindCheckUpdate 检查是否存在新版本
func BindCheckUpdate() bool {
	api := "https://api.lumina.moe/api/gh/repos/project-vrcat/VRChatConfigurationEditor/releases/latest"
	client := http.Client{Timeout: time.Second * 10}
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

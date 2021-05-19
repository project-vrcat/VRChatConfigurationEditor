package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
)

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

func BindReadTextFile(filename string) (content string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	content = string(data)
	return
}

func BindWriteTextFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)
}

func BindSelectDirectory(title string) (dir string, err error) {
	return dialog.Directory().Title(title).Browse()
}

func BindRemoveAll(path string) error {
	return os.RemoveAll(path)
}

func BindAppVersion() string {
	return fmt.Sprintf("%s build-%s", Version, BuildDate)
}

func BindOpen(url string) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	}
	return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
}

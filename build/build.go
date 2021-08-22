//+build script
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	pwd, _ := os.Getwd()
	filename := filepath.Join(pwd, "app/version.go")
	code := fmt.Sprintf(`package app

var (
	// Version 应用版本号
	Version = "%s"
	// GitHash Git Commit Hash
	GitHash = "%s"
)
`, getGitTag(), getGitHash())
	_ = ioutil.WriteFile(filename, []byte(code), 0666)
}

func getGitTag() string {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "v1.0.0"
	}
	return strings.Replace(string(out), "\n", "", -1)
}

func getGitHash() string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.Replace(string(out), "\n", "", -1)
}

//go:generate go-winres simply --icon public/favicon.png
package main

import (
	"embed"
	"log"

	"github.com/project-vrcat/VRChatConfigurationEditor/app"
)

//go:embed public
var publicFiles embed.FS

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	app.PublicFiles = publicFiles
	app.Run()
}

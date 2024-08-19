// Package main ...
package main

import (
	"path/filepath"

	"github.com/assetnote/rod/lib/utils"
)

var slash = filepath.FromSlash

func main() {
	build := utils.S(`// Package assets is generated by "lib/assets/generate"
package assets

// MousePointer for rod
const MousePointer = {{.mousePointer}}

// Monitor for rod
const Monitor = {{.monitor}}

// MonitorPage for rod
const MonitorPage = {{.monitorPage}}
`,
		"mousePointer", get("../../fixtures/mouse-pointer.svg"),
		"monitor", get("monitor.html"),
		"monitorPage", get("monitor-page.html"),
	)

	utils.E(utils.OutputFile(slash("lib/assets/assets.go"), build))
}

func get(path string) string {
	code, err := utils.ReadString(slash("lib/assets/" + path))
	utils.E(err)
	return utils.EscapeGoString(code)
}

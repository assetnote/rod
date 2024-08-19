// Package main ...
package main

import (
	"fmt"

	"github.com/assetnote/rod/lib/launcher"
	"github.com/assetnote/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}

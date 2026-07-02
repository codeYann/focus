// Package main provides the focus CLI command.
package main

import (
	"log"

	"github.com/codeYann/focus/internal/config"
	"github.com/codeYann/focus/internal/tui"
)

func main() {
	error := tui.Run(config.Default())
	if error != nil {
		log.Fatal(error)
	}
}

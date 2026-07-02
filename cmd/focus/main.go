// Package main provides the focus CLI command.
package main

import (
	"log"

	"github.com/codeYann/focus/internal/config"
	"github.com/codeYann/focus/internal/tui"
)

func main() {
	if err := tui.Run(config.Default()); err != nil {
		log.Fatal(err)
	}
}

/*
Copyright © 2025 justin@trahan.dev
*/
package main

import (
	"log/slog"

	"github.com/babbage88/systinker/cmd"
)

func main() {
	configureDefaultLogger(slog.LevelDebug)
	cmd.Execute()
}

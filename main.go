/*
Copyright © 2025 justin@trahan.dev
*/
package main

import (
	"log"
	"log/slog"

	"github.com/babbage88/systinker/cmd"
)

var db *SqliteDb

func main() {
	db = &SqliteDb{Filename: "systinker.db"}
	err := db.initDb()
	if err != nil {
		log.Fatalf("error initializing sqlite db %s", err.Error())
	}
	configureDefaultLogger(slog.LevelDebug)
	cmd.Execute()
}

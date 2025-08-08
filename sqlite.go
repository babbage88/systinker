package main

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var version string

type SqliteDb struct {
	Filename string `json:"fileName"`
	Version  string `json:"version"`
	Db       *sql.DB
}

func (d *SqliteDb) initDb() error {
	file := fmt.Sprintf("file:%s", d.Filename)
	var err error
	d.Db, err = sql.Open("sqlite3", file)
	if err != nil {
		slog.Error("Error opening sqlite db", slog.String("filename", d.Filename), "error", err.Error())
		return err
	}
	err = d.Db.QueryRow(`SELECT sqlite_version()`).Scan(d.Version)
	return err
}

func (d *SqliteDb) ShowVersion() error {
	var err error
	if d.Db == nil {
		err = d.initDb()
	}

	slog.Info("Sqlite", slog.String("Version", d.Version))
	return err
}

func initSqlite(filename string) (*sql.DB, error) {
	file := fmt.Sprintf("file:%s", filename)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		slog.Error("Error opening sqlite db", slog.String("filename", filename), "error", err.Error())
		return nil, err
	}
	err = db.QueryRow(`SELECT sqlite_version()`).Scan(&version)
	return db, err
}

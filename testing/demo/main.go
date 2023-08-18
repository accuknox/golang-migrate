package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/golang-migrate/migrate/v4/internal/cli"
	_ "github.com/golang-migrate/migrate/v4/testing/demo/mongoMigrations"
)

func main() {
	cli.Main("")
}

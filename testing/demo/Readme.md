# Getting started

## Directory Structure

- The following directory structure needs to followed.
```
/
|-->main.go
|-->migrations/
    |--> 000001_add_collection.down.go
    |--> 000001_add_collection.up.go
    |--> 000002_create_user.down.go
    |--> 000002_create_user.up.go
```
## main.go
- In the `main.go` file, import the `golang-migrate` database driver and the `cli` module.
- Import the `migrations` package containing the migration script files.
- Call the `cli.Main()` method in `main`.
```
package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/golang-migrate/migrate/v4/internal/cli"
	_ "github.com/golang-migrate/migrate/v4/testing/demo/mongoMigrations"
)

func main() {
	cli.Main("")
}

```
- Build the custom binary.
```
go build -o migrate-custom *.go
```
## Migrations
- Migration files can be created using the create command.
```
./migrate-custom create -ext go -dir ./mongoMigrations/ -seq add_collection
```
- This will create 2 migration script files, one for up and another for down migration.
- In the migration files, user can define the functions and register them using `migrate.RegisterFunction()`
- To apply migrations:
```
./migrate-custom -database mongodb://127.0.0.1:27017/test?directConnection=true -path ./mongoMigrations/ up
```
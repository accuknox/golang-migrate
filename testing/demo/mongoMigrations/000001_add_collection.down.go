package demo

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
)

func init() {
	migrate.RegisterFunction(dropCollection)
}

func dropCollection(driver database.Driver) error {
	m := driver.(*mongodb.Mongo)
	db := m.GetDbConnection()
	coll := db.Collection("obs_event_1025")
	if err := coll.Drop(context.TODO()); err != nil {
		fmt.Println("Unable to delete collection")
		return err
	}
	fmt.Println("Collection deleted")
	return nil
}

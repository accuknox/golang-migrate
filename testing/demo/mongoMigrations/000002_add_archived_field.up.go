package demo

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	migrate.RegisterFunction(setArchivedFlag)
}

func setArchivedFlag(driver database.Driver) error {
	m := driver.(*mongodb.Mongo)
	db := m.GetDbConnection()
	coll := db.Collection("obs_event_115")
	start := time.Now()
	update := bson.D{{"$set", bson.D{{"archived", false}}}}
	_, err := coll.UpdateMany(context.TODO(), bson.D{{}}, update)
	if err != nil {
		fmt.Printf("setArchivedFlag: Failed to update records. %v\n", err)
		return err
	}
	fmt.Printf("setArchivedFlag: Completed for collection %s. Time taken - %s\n", coll.Name(), time.Since(start))
	return nil
}

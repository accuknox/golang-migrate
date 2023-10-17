package demo

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	migrate.RegisterFunction(addCollection)
}

func addCollection(driver database.Driver) error {
	m := driver.(*mongodb.Mongo)
	db := m.GetDbConnection()
	fmt.Println("Creating Collection...")
	collectionName := "obs_event_1025"
	indexModelCollection := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "cluster_id", Value: 1},
				{Key: "pod_name", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "cluster_id", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "hash_id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{Key: "updated_time_iso", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "cluster_id", Value: 1},
				{Key: "workload.type", Value: 1},
				{Key: "workload.name", Value: 1},
			},
		},
	}
	name, err := db.Collection(collectionName).Indexes().CreateMany(context.TODO(), indexModelCollection)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	fmt.Printf("Collection Created\n", name)
	return nil
}

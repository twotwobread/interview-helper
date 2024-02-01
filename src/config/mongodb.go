package config

import (
	"context"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConfig struct {
	Database struct {
		User     string
		Password string
		Database string
		Port     int
	}
}

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance
var dbConfig DbConfig

func MongodbSetup() error {
	if _, err := toml.DecodeFile("config/config.toml", &dbConfig); err != nil {
		log.Fatal(err)
	}
	mongoURI := fmt.Sprintf("mongodb://%s:%s@localhost:%d/%s",
		dbConfig.Database.User,
		dbConfig.Database.Password,
		dbConfig.Database.Port,
		dbConfig.Database.Database,
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database(dbConfig.Database.Database)
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func GetMongodbInstance() MongoInstance {
	return mg
}

package bootstrap

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInstance struct {
	Client mongo.Client
	Db     mongo.Database
}

func NewDatabaseInstance(env *Env) DatabaseInstance {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@localhost:%d/%s",
		env.Database.User,
		env.Database.Password,
		env.Database.Port,
		env.Database.Database,
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(env.Database.Database)
	di := DatabaseInstance{
		Client: *client,
		Db:     *db,
	}
	return di
}

func CloseMongoDBConnection(client mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

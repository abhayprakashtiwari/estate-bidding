package mongo

import (
	"context"
	"github.com/abhayprakashtiwari/estatebidding/config"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)


var EstateCollection *mongo.Collection

func SetupDb(configuration *config.Configuration)  {
	log.Print("Initializing database")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var Client, err = mongo.Connect(ctx, configuration.Database.ConnectionUri)
	log.Printf("Database uri %s", configuration.Database.ConnectionUri)
	if err != nil {
		log.Fatal(err)
	}
	EstateCollection = Client.Database(configuration.Database.DatabaseName).Collection(configuration.Database.CollectionName)
	log.Print("Database init complete")
}
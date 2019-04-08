package main

import (
	"github.com/abhayprakashtiwari/estatebidding/config"
	"github.com/abhayprakashtiwari/estatebidding/mongo"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/handler"
	"github.com/abhayprakashtiwari/estatebidding/gql"
)

const defaultPort = 8080

func main()  {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var AppConfiguration config.Configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&AppConfiguration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	mongo.SetupDb(&AppConfiguration)

	port := AppConfiguration.Server.Port
	if port == 0 {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}})))
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+ strconv.Itoa(port), nil))
	log.Print("Application up and running")
}

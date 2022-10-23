package main

import (
	"context"
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"log"
	"net/http"
	"os"
	"test-eth/data"
	"test-eth/driver"
	"test-eth/internal/handlers"
)

var portNum = os.Getenv("PORT")
var client *driver.DB
var repo *handlers.Repository

func main() {
	run()
	defer func() {
		if err := client.Mongo.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Starting servers")
	srv := &http.Server{Addr: portNum, Handler: routes()}
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func run() {
	config.AddDriver(yamlv3.Driver)
	err := config.LoadFiles("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
	var uri = config.String("MONGO_URI")

	fmt.Println(uri)
	client, err := driver.ConnectMongoDB(uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)
	go func() { data.InitData(client.Mongo) }()
	fmt.Println(client)
	repo = handlers.NewRepository(client.Mongo)
	listenForBlocks(repo)
	handlers.NewHandlers(repo)
}

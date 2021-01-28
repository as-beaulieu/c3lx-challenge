package main

import (
	"c3lx-challenge/server"
	"c3lx-challenge/service"
	"c3lx-challenge/src/dao"
	"c3lx-challenge/src/logging"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	errors chan error
)

func main() {
	fmt.Println("Starting Challenge API...")

	err := godotenv.Load()
	example := os.Getenv("EXAMPLE")
	fmt.Println(example)

	errors = make(chan error)

	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")

	logger := logging.NewLogger()
	postgres := dao.PostgresBuilder{}.
		SetConnectionString(connectionString).
		Build()
	svc := service.ServiceBuilder{}.
		WithLogger(*logger).
		WithPostgres(postgres).
		Build()

	//http server
	go func() {
		errors <- server.RunHttpServer(svc)
	}()

	err = <-errors
	close(errors)
	if err != nil {
		log.Fatal("error: ", err)
	}
	log.Println("program closed")
}

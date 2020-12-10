package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lucabecci/go-api-blog/internal/data"
	"github.com/lucabecci/go-api-blog/internal/server"
)

func main() {
	port := os.Getenv("PORT")

	serv, err := server.New(port)

	if err != nil {
		log.Fatal(err)
	}

	db := data.New()
	if err := db.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	command := make(chan os.Signal, 1)
	signal.Notify(command, os.Interrupt)
	<-command

	serv.Close()
	data.Close()
}

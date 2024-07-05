package main

import (
	"github.com/joho/godotenv"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"log"
	"time_tracker/configs"
	"time_tracker/internal/app"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title			Time Tracker
//	@version		1.0
//	@description	API for control user's working time

//	@host		localhost:8080
//	@BasePath	/

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.StartService(*config); err != nil {
		log.Fatal(err)
	}
}

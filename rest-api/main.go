package main

import (
	"os"

	"github.com/grantgariepy/rest-api/database"
	"github.com/joho/godotenv"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer close databse
	defer database.CloseMongoDB()

	app := generateApp()

	// get the port from env

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func initApp() error {
	// setup env
	err := loadENV()
	if err != nil {
		return err
	}
	// setup databse
	err = database.StartMongoDB()

	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

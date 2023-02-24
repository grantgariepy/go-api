package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/grantgariepy/rest-api/database"
	"github.com/grantgariepy/rest-api/handlers"
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

func generateApp() *fiber.App {
	app := fiber.New()

	// create health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create library group and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.GetLibraries)
	libGroup.Post("/", handlers.CreateLibrary)
	libGroup.Post("/book", handlers.CreateBook)

	return app

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

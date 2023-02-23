package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/grantgariepy/rest-api/database"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer close databse
	defer database.CloseMongoDB()

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		//write a random title to db
		sampleDoc := bson.M{"name": "sample todo"}
		collection := database.GetCollection("todos")
		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error Inserting Todo")
		}

		// send down info about todo
		return c.JSON(nDoc)
	})

	app.Listen(":3000")
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
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/grantgariepy/rest-api/database"
	"github.com/grantgariepy/rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exists" bson:"books"`
}

// GET
func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}
	var libraries []models.Library
	if err := cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}
	return c.JSON(libraries)
}

// POST
func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := new(libraryDTO)

	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}
	// mongodb question
	nLibrary.Empty = make([]string, 0)

	libraryCollection := database.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}

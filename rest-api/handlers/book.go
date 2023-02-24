package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/grantgariepy/rest-api/database"
	"github.com/grantgariepy/rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

type newBookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryId string `json:"libraryId" bson:"libraryId"`
}

func CreateBook(c *fiber.Ctx) error {
	createData := new(newBookDTO)

	if err := c.BodyParser(createData); err != nil {
		return err
	}

	// get collection reference
	coll := database.GetCollection("libraries")

	// get the filter
	filter := bson.D{{Key: "id", Value: createData.LibraryId}}
	nBookData := models.Book{
		Title:  createData.Title,
		Author: createData.Author,
		ISBN:   createData.ISBN,
	}
	updatePayload := bson.M{"$push": bson.M{"books": nBookData}}

	// update the library
	res, err := coll.UpdateOne(context.TODO(), filter, updatePayload)
	if err != nil {
		return err
	}

	fmt.Println(res.UpsertedID)

	return c.SendString("Book Created Successfully")
}

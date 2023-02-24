package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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

	fmt.Println(createData)
	return nil
}

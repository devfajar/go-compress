package main

import (
	"github.com/gofiber/fiber/v2"
	"io"
)

func main() {
	app := fiber.New()

	app.Static("/uploads", "./uploads")

	app.Post("/", func(c *fiber.Ctx) error {
		fileheader, err := c.FormFile("picture")
		if err != nil {
			panic(err)
		}

		file, err := fileheader.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()

		buffer, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		errDir := createFolder("uploads")
		if errDir != nil {
			panic(errDir)
		}

		filename, err := imageProcessing(buffer, 40, "uploads")
		if err != nil {
			panic(err)
		}
		return c.JSON(fiber.Map{
			"picture": "http://localhost:3000/uploads/" + filename,
		})
	})

	app.Listen(":3000")
}

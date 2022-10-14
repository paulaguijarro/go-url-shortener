package server

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/paulaguijarro/go-url-shortener/database"
	"github.com/paulaguijarro/go-url-shortener/utils"
)

func getAllGoShorts(ctx *fiber.Ctx) error {
	goshorts, err := database.GetAllGoShorts()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error getting all GoShort links " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(goshorts)
}

func getGoShort(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}
	goshort, err := database.GetGoShort(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error could not retreive goshort from DB " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(goshort)
}

func createGoShort(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var goshort database.GoShort
	err := ctx.BodyParser(&goshort)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	// Validate URL
	_, err = url.ParseRequestURI(goshort.Redirect)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error validating url " + err.Error(),
		})
	}

	if goshort.Random {
		goshort.Goshort = utils.RandomURL(8)
	}

	goshort, err = database.CreateGoShort(goshort)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error creating goshort in DB " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(goshort)

}

func updateGoshort(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var goshort database.GoShort
	err := ctx.BodyParser(&goshort)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	goshort, err = database.UpdateGoShort(goshort)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error updating goshort in DB " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(goshort)
}

func deleteGoshort(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing id " + err.Error(),
		})
	}

	err = database.DeleteGoShort(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error deleting goshort from DB " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goshort has been deleted",
	})
}

func redirect(ctx *fiber.Ctx) error {
	goshortUrl := ctx.Params("redirect")

	goshort, err := database.FindByGoShortUrl(goshortUrl)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "could not find goshort in DB " + err.Error(),
		})
	}

	// Stats: update times url clicked
	goshort.Clicked += 1
	_, err = database.UpdateGoShort(goshort)
	if err != nil {
		fmt.Printf("error updating stats: %v\n", err)
	}

	return ctx.Redirect(goshort.Redirect, fiber.StatusTemporaryRedirect)
}

func RunServer() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/goshort", getAllGoShorts)
	app.Get("/goshort/:id", getGoShort)
	app.Post("/goshort", createGoShort)
	app.Patch("/goshort", updateGoshort)
	app.Delete("/goshort/:id", deleteGoshort)

	app.Get("/r/:redirect", redirect)

	app.Listen(":3000")
}

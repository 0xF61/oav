package main

import (
	"log"

	"github.com/AlienVault-OTX/OTX-Go-SDK/src/otxapi"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var client *otxapi.Client

func main() {
	// Configure oav client
	// os.Setenv("X_OTX_API_KEY", "<ChangeMe>")
	client = otxapi.NewClient(nil)

	// Handle Web Reqs
	engine := html.New("./views/", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", indexHandler)
	app.Get("/info", infoHandler)

	log.Fatalln(app.Listen(":3000"))
}

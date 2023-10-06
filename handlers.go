package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AlienVault-OTX/OTX-Go-SDK/src/otxapi"
	"github.com/gofiber/fiber/v2"
)

var user_detail *otxapi.UserDetail

func indexHandler(c *fiber.Ctx) error {

	p := c.Query("page", "1")
	pg, err := strconv.Atoi(p)
	if err != nil {
		pg = 1
	}

	opt := &otxapi.ListOptions{Page: pg, PerPage: 10}
	ti, _, err := client.ThreatIntel.List(opt)
	if err != nil {
		log.Println("Feed Error:", err)
	}

	for _, pls := range ti.Pulses {

		for i, ind := range pls.Indicators {
			if !strings.Contains(*ind.Type, "FileHash") {
				ind.Type = nil
				pls.Indicators[i] = ind
			}
		}

	}

	return c.Render("threat", ti)
}

func infoHandler(c *fiber.Ctx) error {

	if c.Query("force") != "" || user_detail == nil {
		log.Println("Update the user_detail")
		ud, _, err := client.UserDetail.Get()
		user_detail = &ud
		if err != nil {
			fmt.Printf("Fatal error: %v\n\n", err)
			fmt.Println("Please ensure your API KEY is valid (i.e. `echo $X_OTX_API_KEY`)")
			os.Exit(1)
		}
	}

	return c.Render("info", user_detail)
}

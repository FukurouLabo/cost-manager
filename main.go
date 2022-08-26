package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func testTicketList() []map[string]string {
	ticketList := []map[string]string{
		{
			"name":   "test ticket",
			"url":    "https://test.cost-manager.com",
			"number": "CV-99999",
		},
		{
			"name":   "test ticket 2",
			"url":    "https://test.cost-manager.com",
			"number": "CV-99999",
		},
	}
	return ticketList
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "cost-manager",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(testTicketList)
	app.Run()
}

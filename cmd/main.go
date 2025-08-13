package main

import (
	"bambah-desktop/internal/ui"
	"log"

	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myApp.SetMetadata(getAppMetadata())
	
	mainWindow := ui.NewMainWindow(myApp)
	
	if err := mainWindow.ShowAndRun(); err != nil {
		log.Fatal(err)
	}
}

func getAppMetadata() *app.Metadata {
	return &app.Metadata{
		ID:          "io.bambah.desktop",
		Name:        "Bambah",
		Description: "Personal Finance Management",
		Version:     "1.0.0",
	}
}
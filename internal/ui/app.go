package ui

import (
	"bambah-desktop/internal/ui/forms"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/marcelocg/bambah-sdk"
)

type MainWindow struct {
	app     fyne.App
	window  fyne.Window
	service bambah.FinancialService
}

func NewMainWindow(app fyne.App) *MainWindow {
	service, err := bambah.NewFinancialServiceFromEnv()
	if err != nil {
		log.Fatalf("Failed to create financial service: %v", err)
	}

	window := app.NewWindow("Bambah - Controle Financeiro")
	window.Resize(fyne.NewSize(600, 500))
	window.SetIcon(theme.DocumentIcon())

	return &MainWindow{
		app:     app,
		window:  window,
		service: service,
	}
}

func (mw *MainWindow) ShowAndRun() error {
	content := mw.buildContent()
	mw.window.SetContent(content)
	mw.window.ShowAndRun()
	return nil
}

func (mw *MainWindow) buildContent() *container.Border {
	title := widget.NewLabelWithStyle("Bambah - Controle Financeiro", 
		fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	entryForm := forms.NewEntryForm(mw.service)

	content := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator()),
		nil,
		nil,
		nil,
		entryForm.GetContainer(),
	)

	return content
}
package main

import (
	"testeditor/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("com.testeditor")
	w := a.NewWindow("Testeditor")
	mainContent := ui.CreateMainContent(w, a.Storage())

	w.SetContent(mainContent.MakeUI())
	w.Resize(fyne.NewSize(600, 600))

	w.ShowAndRun()
}

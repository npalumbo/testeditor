package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"testeditor/internal/ui"
)

func main() {
	a := app.NewWithID("com.testeditor")
	w := a.NewWindow("Testeditor")

	mainContent := ui.CreateMainContentSpecific(w)
	w.SetContent(mainContent.MakeUI())

	//mainContent := ui.CreateMainContent(w)
	//w.SetContent(mainContent.MakeUI())

	w.Resize(fyne.NewSize(600, 600))

	w.ShowAndRun()
}

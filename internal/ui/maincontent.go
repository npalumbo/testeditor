package ui

import (
	"fmt"
	testeditorText "testeditor/internal/text"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainContent struct {
	container *fyne.Container
	Text      *widget.TextGrid
	DebugLog  *widget.Label
	cursor    int
	//
	realText testeditorText.TextBackend
}

func (m *MainContent) MakeUI() fyne.CanvasObject {
	return m.container
}

func CreateMainContent(parent fyne.Window, stor fyne.Storage) MainContent {

	debugLog := widget.NewLabel("Cursor: 0")
	text := widget.NewTextGridFromString("Edit me")
	currentContainer := container.NewBorder(nil, debugLog, nil, nil, text)

	currentContainer.Add(text)
	mainContent := MainContent{container: currentContainer,
		Text:     text,
		DebugLog: debugLog,
		cursor:   0,
		realText: testeditorText.CreateInternalText(),
	}

	parent.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch key := ke.Name; key {
		case fyne.KeyBackspace:
			if mainContent.cursor > 0 {
				mainContent.realText.Delete(mainContent.cursor-1, mainContent.cursor)
				mainContent.cursor--
			}
		case fyne.KeySpace:
			mainContent.realText.Insert(mainContent.cursor, " ")
			mainContent.cursor++
		default:
			mainContent.realText.Insert(mainContent.cursor, string(ke.Name))
			mainContent.cursor++
		}
		// fmt.Println(mainContent.realText.Render())
		mainContent.Text.SetText(mainContent.realText.Render())
		mainContent.Text.Refresh()
		debugLog.SetText("Cursor: " + fmt.Sprint(mainContent.cursor))

	})

	return mainContent

}

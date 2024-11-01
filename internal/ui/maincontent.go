package ui

import (
	cursor2 "testeditor/internal/cursor"
	testeditorText "testeditor/internal/text"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainContent struct {
	mainWindow fyne.Window
	container  *fyne.Container
	Text       *widget.TextGrid
	cursor     *cursor2.Cursor
	//
	textBackend testeditorText.TextBackend
}

func (m *MainContent) MakeUI() fyne.CanvasObject {
	return m.container
}

func CreateMainContent(parent fyne.Window) MainContent {

	cursor := cursor2.CreateCursor()
	text := widget.NewTextGridFromString("Edit me")
	currentContainer := container.NewBorder(nil, cursor.DebugLabel, nil, nil, text)

	currentContainer.Add(text)
	mainContent := MainContent{
		mainWindow:  parent,
		container:   currentContainer,
		Text:        text,
		cursor:      cursor,
		textBackend: testeditorText.CreateInternalText(),
	}

	parent.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch key := ke.Name; key {
		case fyne.KeyEscape:
			mainContent.mainWindow.Close()
		case fyne.KeyBackspace:
			if mainContent.cursor.CanBackspace() {
				mainContent.textBackend.Delete(mainContent.cursor.CurrentPosition()-1, mainContent.cursor.CurrentPosition())
				mainContent.cursor.Dec()
			}
		case fyne.KeySpace:
			mainContent.textBackend.Insert(mainContent.cursor.CurrentPosition(), " ")
			mainContent.cursor.Inc()
		default:
			mainContent.textBackend.Insert(mainContent.cursor.CurrentPosition(), string(ke.Name))
			mainContent.cursor.Inc()
		}
		mainContent.Text.SetText(mainContent.textBackend.Render())
		mainContent.Text.Refresh()
		cursor.DebugLabel.Refresh()
	})

	return mainContent

}

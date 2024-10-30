package ui

import (
	cursor2 "testeditor/internal/cursor"
	testeditorText "testeditor/internal/textspecific"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainContentSpecific struct {
	mainWindow fyne.Window
	container  *fyne.Container
	Text       *widget.TextGrid
	cursor     *cursor2.Cursor
	//
	textBackend testeditorText.TextBackendSpecific
}

func (m *MainContentSpecific) MakeUI() fyne.CanvasObject {
	return m.container
}

func CreateMainContentSpecific(parent fyne.Window) MainContentSpecific {

	cursor := cursor2.CreateCursor()
	text := widget.NewTextGridFromString("Edit me")
	currentContainer := container.NewBorder(nil, cursor.DebugLabel, nil, nil, text)

	currentContainer.Add(text)
	mainContent := MainContentSpecific{
		mainWindow:  parent,
		container:   currentContainer,
		Text:        text,
		cursor:      cursor,
		textBackend: testeditorText.CreateInternalTextSpecific(),
	}

	parent.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch key := ke.Name; key {
		case fyne.KeyEscape:
			mainContent.mainWindow.Close()
		case fyne.KeyBackspace:
			if mainContent.cursor.CanBackspace() {
				mainContent.textBackend.Backspace(*mainContent.cursor)
				mainContent.cursor.Dec()
			}
		case fyne.KeySpace:
			mainContent.textBackend.Insert(*mainContent.cursor, " ")
			mainContent.cursor.Inc()
		default:
			mainContent.textBackend.Insert(*mainContent.cursor, string(ke.Name))
			mainContent.cursor.Inc()
		}
		mainContent.Text.SetText(mainContent.textBackend.Render())
		mainContent.Text.Refresh()
		cursor.DebugLabel.Refresh()
	})

	return mainContent

}

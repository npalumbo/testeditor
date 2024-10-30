package textspecific

import (
	"testeditor/internal/cursor"
)

type TextBackendSpecific interface {
	Insert(cursor cursor.Cursor, newText string)
	Backspace(cursor cursor.Cursor)
	Render() string
}

type internalTextStorageSpecific struct {
	rawText string
}

func CreateInternalTextSpecific() TextBackendSpecific {
	return &internalTextStorageSpecific{rawText: ""}
}

func (i *internalTextStorageSpecific) Render() string {
	return i.rawText
}

func (i *internalTextStorageSpecific) Insert(cursor cursor.Cursor, newText string) {
	position := cursor.CurrentPosition()
	if position < 0 || position > len(i.rawText) {
		return
	}

	prefix := i.rawText[:position]
	suffix := i.rawText[position:]

	newString := prefix + newText + suffix
	i.rawText = newString
}

func (i *internalTextStorageSpecific) Backspace(cursor cursor.Cursor) {
	start := cursor.CurrentPosition() - 1
	end := cursor.CurrentPosition()
	if start < 0 || end > len(i.rawText) || start > end {
		return
	}

	i.rawText = i.rawText[:start] + i.rawText[end:]
}

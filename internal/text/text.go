package text

type TextBackend interface {
	Insert(position int, newText string)
	Delete(start int, end int)
	Render() string
}

type internalTextStorage struct {
	rawText string
}

func CreateInternalText() TextBackend {
	return &internalTextStorage{rawText: ""}
}

func (i *internalTextStorage) Render() string {
	return i.rawText
}

func (i *internalTextStorage) Insert(position int, newText string) {
	if position < 0 || position > len(i.rawText) {
		return
	}

	prefix := i.rawText[:position]
	suffix := i.rawText[position:]

	newString := prefix + newText + suffix
	i.rawText = newString
}

func (i *internalTextStorage) Delete(start, end int) {
	if start < 0 || end > len(i.rawText) || start > end {
		return
	}

	i.rawText = i.rawText[:start] + i.rawText[end:]
}

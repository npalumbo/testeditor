package text

import "strings"

type TextBackendExtension interface {
	TextBackend
	SearchAndReplace(search, replace string)
}

func CreateInternalTextExtension() TextBackendExtension {
	return &internalTextStorage{rawText: ""}
}

func (i *internalTextStorage) SearchAndReplace(search, replace string) {
	from, to, found := i.findSubstringRange(search)
	if found {
		i.Delete(from, to)
		i.Insert(from, replace)
	}
}

func (i *internalTextStorage) findSubstringRange(substr string) (int, int, bool) {
	start := strings.Index(i.rawText, substr)
	if start == -1 {
		return 0, 0, false // Substring not found
	}
	end := start + len(substr)
	return start, end, true
}

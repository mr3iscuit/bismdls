package types

type TextDocumentItem struct {
	URI        string
	LanguageId string
	Version    int
	Text       string
}

func NewTextDocumentItem(uri, languageId string, version int, text string) *TextDocumentItem {
	return &TextDocumentItem{
		URI:        uri,
		LanguageId: languageId,
		Version:    version,
		Text:       text,
	}
}

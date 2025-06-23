package text_document

import (
	"github.com/mr3iscuit/bismdls/lsp"
)

type DidChangeDocumentNotification struct {
	lsp.Notification
	Params DidChangeDocumentParams `json:"params"`
}

type DidChangeDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextCodumentContentChangeEvent `json:"contentChanges"`
}

type TextCodumentContentChangeEvent struct {
	Text string `json:"text"`
}

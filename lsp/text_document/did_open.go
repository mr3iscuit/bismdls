package text_document

import "github.com/mr3iscuit/bismdls/lsp"

type DidOpenTextDocumentNotification struct {
	lsp.Notification
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	TextDocument   TextDocumentItem                 `json:"textDocument"`
	ContentChanges []TextCodumentContentChangeEvent `json:"contentChanges"`
}

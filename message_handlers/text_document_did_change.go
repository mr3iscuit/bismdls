package message_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/mr3iscuit/bismdls/analysis"
	"github.com/mr3iscuit/bismdls/lsp/text_document"
)

type TextDocumentDidChange struct {
	HandlerStrategy
}

func (td *TextDocumentDidChange) Handle(msg []byte) (any, error) {
	var request text_document.DidChangeDocumentNotification
	if err := json.Unmarshal(msg, &request); err != nil {
		return nil, fmt.Errorf("could not unmarshal the message: %s", string(msg))
	}

	state := analysis.NewStateInstance()
	for _, change := range request.Params.ContentChanges {
		state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
	}

	return nil, nil
}

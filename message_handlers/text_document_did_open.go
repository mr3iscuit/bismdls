package message_handlers

import (
	"encoding/json"

	"github.com/mr3iscuit/bismdls/analysis"
	"github.com/mr3iscuit/bismdls/logger"
	"github.com/mr3iscuit/bismdls/lsp/text_document"
)

type TextDocumentDidOpen struct {
	HandlerStrategy
}

func (td *TextDocumentDidOpen) Handle(msg []byte) (any, error) {
	var request text_document.DidOpenTextDocumentNotification

	if err := json.Unmarshal(msg, &request); err != nil {
		logger.GetLogger().Printf("could not unmarshal the message: %s", string(msg))
	}

	logger.GetLogger().Printf("Text document opened: {uri: %s, text: %s}",
		request.Params.TextDocument.URI, request.Params.TextDocument.Text,
	)

	state := analysis.NewStateInstance()
	state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	return nil, nil
}

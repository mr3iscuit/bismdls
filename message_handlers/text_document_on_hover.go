package message_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/mr3iscuit/bismdls/lsp"
	"github.com/mr3iscuit/bismdls/lsp/text_document"
)

type TextDocumentOnHover struct {
	HandlerStrategy
}

func (td *TextDocumentOnHover) Handle(msg []byte) (any, error) {
	var request text_document.HoverRequest
	if err := json.Unmarshal(msg, &request); err != nil {
		return nil, fmt.Errorf("could not unmarshal the message: %s", string(msg))
	}

	response := text_document.HoverResponse{
		Response: lsp.Response{
			ID:  &request.ID,
			RPC: "2.0",
		},
		Result: text_document.HoverResult{
			Contents: "Hello from lsp server",
		},
	}

	return response, nil
}

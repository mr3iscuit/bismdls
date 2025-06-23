package message_handlers

import "github.com/mr3iscuit/bismdls/logger"

type TextDocumentDidClose struct {
	HandlerStrategy
}

func (td *TextDocumentDidClose) Handle(msg []byte) (any, error) {
	logger.GetLogger().Printf("Text document closed: %s", string(msg))

	return nil, nil
}

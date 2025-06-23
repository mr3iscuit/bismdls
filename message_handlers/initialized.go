package message_handlers

import "github.com/mr3iscuit/bismdls/logger"

type Initialized struct {
	HandlerStrategy
}

func (initialized *Initialized) Handle(msg []byte) (any, error) {
	logger.GetLogger().Printf("Initialized notification received: %s", string(msg))
	logger.GetLogger().Printf("LSP server is now fully initialized and ready!")

	return nil, nil
}

package handler

import (
	"encoding/json"
	"errors"

	"github.com/mr3iscuit/bismdls/logger"
	"github.com/mr3iscuit/bismdls/lsp"
)

type Initialize struct {
	HandlerStrategy
}

func (initalize *Initialize) Handle(msg []byte) (any, error) {
	logger.GetLogger().Printf("Initialize handler called with message: %s", string(msg))

	var request lsp.InitializeRequest

	if err := json.Unmarshal(msg, &request); err != nil {
		return errors.New("could not unmarshal the message: " + string(msg)), nil
	}

	logger.GetLogger().Printf("ClientInfo: %+v", request.Params.ClientInfo)
	logger.GetLogger().Printf("I am connected to %s, %s.\n", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

	// Send initialize response back to Neovim
	response := lsp.InitializeResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &request.ID,
		},
		Result: lsp.InitializeResponseResult{
			Capabilities: lsp.ServerCapabilities{
				TextDocumentSync: 1,
				HoverProvider:    true,
			},
			ServerInfo: lsp.ServerInfo{
				Name:    "bisgols",
				Version: "0.1.0",
			},
		},
	}

	return response, nil
}

package text_document

import "github.com/mr3iscuit/bismdls/lsp"

type DefinitionRequest struct {
	lsp.Request
	Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
	TextDocumentPositionParams
}

type DefinitionResponse struct {
	lsp.Response
	Result Location `json:"result"`
}

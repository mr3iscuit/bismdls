package text_document

import "github.com/mr3iscuit/bismdls/lsp"

type HoverRequest struct {
	lsp.Request
	Params HoverParams `json:"params"`
}

type HoverParams struct {
	TextDocumentPositionParams
}

type HoverResponse struct {
	lsp.Response
	Result HoverResult `json:"result"`
}

type HoverResult struct {
	Contents string `json:"contents"`
}

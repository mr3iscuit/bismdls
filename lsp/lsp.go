package lsp

type Request struct {
	RPC string `json:"jsonrpc"`
	ID  int    `json:"id"`
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	ID     *int   `json:"id"`
	Method string `json:"method"`
}

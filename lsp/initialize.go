package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResponseResult `json:"result"`
}

type InitializeResponseResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
    TextDocumentSync int `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResponseResult{
			Capabilities: ServerCapabilities{
                TextDocumentSync: 1,
            },
			ServerInfo: ServerInfo{
				Name:    "edulsp",
				Version: "0.0.1",
			},
		},
	}
}

package lsp

// See: https://microsoft.github.io/language-server-protocol/specifications/base/0.9/specification/#contentPart
type Request struct {
    RPC string `json:"jsonrpc"`
    ID int `json:"id"`
    Method string `json:"method"`

    // We'll handle params later
    // See: lsp.InitializeRequestParams
}

type Response struct {
    RPC string `json:"jsonrpc"`
    ID *int `json:"id,omitempty"`

    // Result
    // Error
}

type Notification struct {
    Method string `json:"method"`

    // We'll handle params later
}


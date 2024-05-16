package lsp

type Request struct {
    RPC string `json:"jsonrpc"`
    ID int `json:"id"`
    Method string `json:"method"`

    // We'll handle params later
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


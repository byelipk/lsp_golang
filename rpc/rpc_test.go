package rpc_test

import "testing"
import "edu_lsp_go/rpc"

type EncodingExample struct {
    Testing bool
}

func TestEncodeMessage(t *testing.T) {
    expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
    actual := rpc.EncodeMessage(EncodingExample{Testing: true})

    if expected != actual {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestDecodeMessage(t *testing.T) {
    incoming := "Content-Length: 34\r\n\r\n{\"Testing\":true, \"method\": \"test\"}"
    method, content, err := rpc.DecodeMessage([]byte(incoming))

    contentLen := len(content)

    _ = method

    if err != nil {
        t.Fatalf("Error decoding message: %s", err)
    }

    if contentLen != 34 {
        t.Fatalf("Expected content length 16, got %d", contentLen)
    }
}

// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
    return 0, nil, nil
}

package rpc

import (
	"encoding/json"
	"fmt"
    "bytes"
    "errors"
    "strconv"
)

type BaseMessage struct {
    Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	// Read the "Content-Length" field to see how many bytes
	// the message contains.
	headerSeparator := "\r\n"
	contentSeparator := "\r\n"
	contentLen := len(content)

	return fmt.Sprintf(
		"Content-Length: %d%s%s%s",
		contentLen,
		headerSeparator,
		contentSeparator,
		content)

}

func DecodeMessage(msg []byte) (string, []byte, error) {
    header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

    if !found {
        return "", nil, errors.New("Did not find separator")
    }

    contentLenBytes := header[len("Content-Length: "):]
    contentLen, err := strconv.Atoi(string(contentLenBytes))

    if err != nil {
        return "", nil, err
    }

    var baseMessage BaseMessage
    if err := json.Unmarshal(content[:contentLen], &baseMessage); err != nil {
        return "", nil, err
    }

    return baseMessage.Method, content, nil
}

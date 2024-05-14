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

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
    header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

    // If we don't have the full header, wait.
    if !found {
        return 0, nil, nil
    }

    contentLenBytes := header[len("Content-Length: "):]
    contentLen, err := strconv.Atoi(string(contentLenBytes))
    // If we can't parse the content length, return an error.
    if err != nil {
        return 0, nil, err
    }

    // Wait until we have the full message.
    if len(content) < contentLen {
        return 0, nil, nil
    }

    // Return the full message.
    totalLen := len(header) + 4 + contentLen
    return totalLen, data[:contentLen], nil
}

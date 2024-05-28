package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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
		content,
	)

}

func DecodeMessage(msg []byte) (method string, contentLength int, content []byte, error error) {
	separatorLength := 4
	separator := []byte{'\r', '\n', '\r', '\n'}
	headerEndIdx := bytes.Index(msg, separator)

	if headerEndIdx == -1 {
		return "", 0, nil, errors.New("[Decoder] Did not find separator")
	}

	headers := msg[:headerEndIdx]
	content = msg[headerEndIdx+separatorLength:]

	headerLines := strings.Split(string(headers), "\r\n")
	contentLength = -1

	for _, h := range headerLines {
		parts := strings.SplitN(h, ": ", 2)

		if len(parts) != 2 {
			continue
		}

		if parts[0] == "Content-Length" {
			length, err := strconv.Atoi(parts[1])

			if err != nil {
				return "", 0, nil, err
			}

			contentLength = length
		}
	}

	if contentLength == -1 {
		return "", 0, nil, errors.New("[Decoder] No content length found")
	}

	if len(content) < contentLength {
		actualLength := len(content)
		errorMessage := fmt.Sprintf("[Decoder] Content length is less than expected. Actual: %d, Expected: %d", actualLength, contentLength)
		return "", 0, nil, errors.New(errorMessage)
	}

	content = content[:contentLength]

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", 0, nil, err
	}

	return baseMessage.Method, contentLength, content, nil
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	separator := []byte{'\r', '\n', '\r', '\n'}
	headerEndIdx := bytes.Index(data, separator)

	if headerEndIdx == -1 {
		if atEOF {
			return len(data), nil, nil
		}
		return 0, nil, nil
	}

	headers := data[:headerEndIdx]
	headerLines := bytes.Split(headers, []byte("\r\n"))
	contentLength := -1

	for _, h := range headerLines {
		parts := bytes.SplitN(h, []byte(": "), 2)

		if len(parts) != 2 {
			continue
		}

		if bytes.Equal(parts[0], []byte("Content-Length")) {
			length, err := strconv.Atoi(string(parts[1]))

			if err != nil {
				return 0, nil, err
			}

			contentLength = length
		}
	}

	if contentLength == -1 {
		return 0, nil, errors.New("[Split] No content length found")
	}

	totalLength := headerEndIdx + len(separator) + contentLength

	// Tell the scanner to read more data because we don't
	// have a full message yet.
	if len(data) < totalLength {
		return 0, nil, nil
	}

	// Return the full message.
	return totalLength, data[:totalLength], nil
}

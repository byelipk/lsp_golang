package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "encoding/json"
    "path/filepath"

	"edu_lsp_go/rpc"
	"edu_lsp_go/lsp"
)

func main() {
	fmt.Println("Starting LSP server...")

    loggerPath := filepath.Join(os.Getenv("XDG_STATE_HOME"), "nvim", "edulsp.log")
	logger := getLogger(loggerPath)
	logger.Println("Starting up!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contentLength, contentBody, err := rpc.DecodeMessage(msg)

        logger.Println("Bytes received: ", len(msg))

		if err != nil {
			logger.Println("Error decoding message: ", err)
			continue
		}

		handleMessage(logger, method, contentLength, contentBody)
	}

}

func handleMessage(logger *log.Logger, method string, contentLength int, contentBody []byte) {
    logger.Println("-----")
	logger.Println("Received message with method: ", method)
    logger.Println("Content length: ", contentLength)
    logger.Println("Content body: ", string(contentBody))
    logger.Println("-----")

    switch method {
    case "initialize":
        var request lsp.InitializeRequest
        if err := json.Unmarshal(contentBody, &request); err != nil {
            logger.Println("Error unmarshalling initialize request: ", err)
        }
        logger.Println("Connected with client: ", request.Params.ClientInfo.Name)
        logger.Println("Client version: ", request.Params.ClientInfo.Version)
    default:
        logger.Println("Method not implemented: ", method)
    }
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(
		filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file")
	}

	return log.New(logfile, "[edulsp]", log.LstdFlags)
}

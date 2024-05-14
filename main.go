package main

import (
    "fmt"
    "bufio"
    "os"
    "log"

    "edu_lsp_go/rpc"
)

func main() {
    fmt.Println("Starting LSP server...")

    logger := getLogger("./edulsp.log")
    logger.Println("Starting up!")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    for scanner.Scan() {
        msg := scanner.Text()
        handleMessage(logger, msg)
    }

}

func handleMessage(logger *log.Logger, msg any) {
    fmt.Println("Received message: ", msg)
    logger.Println("Received message: ", msg)
}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic("Failed to open log file")
    }

    return log.New(logfile, "[edulsp]", log.LstdFlags)
}

package main

import (
	"bufio"
	"os"

	"github.com/mr3iscuit/bismdls/handler"
	"github.com/mr3iscuit/bismdls/logger"
	"github.com/mr3iscuit/bismdls/rpc"
)

func main() {
	logger := logger.GetLogger()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Some kind of error happened: %s", err)
		}

		logger.Printf("GotMessageType : [%s]", method)

		handlerRegistry := handler.HandlerRegistryInstannce()
		handler := handlerRegistry.GetHandler(method)
		if handler != nil {
			response, err := handler.Handle(contents)
			if err != nil {
				logger.Printf("Error handling message for method [%s]: %s", method, err)
			}
			writer := os.Stdout

			if response != nil {
				writer.Write(rpc.EncodeMessage(response))
			}
		} else {
			logger.Printf("No handler registered for method: %s", method)
		}
	}
}

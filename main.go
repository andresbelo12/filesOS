package main

import (
	"fmt"

	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
	"github.com/andresbelo12/fileOS/handler"
)

const (
	LOCALHOST   = "127.0.0.1"
	SERVER_PORT = "8080"
)

func main() {

	logSystem := handler.CreateLogSystem()
	listener := handler.CreateListener()

	err := logSystem.InitLogSystem()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	connection := connectToServer(&logSystem)
	kernelHandler.ListenServer(&logSystem, listener, &connection)

}

func connectToServer(logSystem *handler.LogSystem) (connection kernelModel.ClientConnection) {
	connection = kernelModel.ClientConnection{ServerHost: LOCALHOST, ServerPort: SERVER_PORT}
	firstMessage := kernelModel.Message{
		Source: kernelModel.MD_FILES, 
		Destination: kernelModel.MD_KERNEL, 
		Command: kernelModel.CMD_START, 
		Message: "listening",
	}

	if err := kernelHandler.EstablishClient(&connection, firstMessage); err != nil {
		fmt.Println(err)
		return
	}

	return
}

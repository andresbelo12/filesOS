package main

import (
	"bufio"
	"fmt"
	"os"

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

	err := logSystem.InitLogSystem()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	connection := connectToServer(&logSystem)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Print(text)

		a := kernelModel.Message{Source: "name", Command: "nasdkasda"}
		//fmt.Fprintf(connection, text+"\n")
		(*connection.ServerConnection).Write(a.ToJson())

		/*message := handler.ReadMessage(&connection)
		fmt.Println(message)*/

	}

}

func connectToServer(logSystem *handler.LogSystem) (connection kernelModel.ClientConnection) {
	connection = kernelModel.ClientConnection{ServerHost: LOCALHOST, ServerPort: SERVER_PORT}
	if err := kernelHandler.EstablishClient(&connection, kernelModel.Message{Source: kernelModel.MD_FILES, Destination: "KERNEL", Command: kernelModel.CMD_START, Message: "Listening"}); err != nil {
		fmt.Println(err)
		return
	}

	listener := handler.CreateListener()

	go kernelHandler.ListenServer(logSystem, listener, &connection)
	return
}

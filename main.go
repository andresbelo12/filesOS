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
	MODULE      = "FILES"
	CMD_START   = "start"
	CMD_SEND    = "send"
	CMD_INFO    = "info"
	CMD_STOP    = "stop"
)

func main() {
	logSystem := handler.CreateLogSystem()
	err := logSystem.InitLogSystem()
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	/*for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Print(text)

		a := model.Message{Source: "name", Command: "nasdkasda"}
		//fmt.Fprintf(connection, text+"\n")
		(*connection.ServerConnection).Write(a.ToJson())

		/*message := handler.ReadMessage(&connection)
		fmt.Println(message)

	}*/

}

func connectToServer(connection kernelModel.ClientConnection) {
	connection = kernelModel.ClientConnection{ServerHost: LOCALHOST, ServerPort: SERVER_PORT}
	if err := kernelHandler.EstablishClient(&connection, kernelModel.Message{Source: MODULE, Destination: "KERNEL", Command: CMD_START, Message: "Listening"}); err != nil {
		fmt.Println(err)
		return
	}

	go kernelHandler.ListenServer(&connection)
	return
}

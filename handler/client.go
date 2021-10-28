package handler

import (
	"fmt"
	"strings"

	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
	"github.com/andresbelo12/fileOS/model"
)

type ClientListener struct {
	LogSystem *LogSystem
}

func CreateListener(logSystem *LogSystem) kernelModel.CommunicationListener {
	return ClientListener{LogSystem: logSystem}
}

func (listener ClientListener) ProcessMessage(connection interface{}, message *kernelModel.Message) (err error) {

	if err = listener.LogSystem.WriteLog(message.Destination, model.MessageToLog(message)); err != nil {
		fmt.Println(err)
	}

	if err = listener.LogSystem.WriteLog(message.Source, model.MessageToLog(message)); err != nil {
		fmt.Println(err)
	}

	messageBody := strings.Split(message.Message, ":")
	conn := connection.(**kernelModel.ClientConnection)

	if len(messageBody) == 2 {
		if messageBody[0] == "create" {
			err = listener.ActionCreateFolder(*conn, message)
			return
		}
	
		if messageBody[0] == "delete" {
			err = listener.ActionDeleteFolder(*conn, message)
			return
		}
		
		if messageBody[0] == "log"{
			fmt.Println(message)
			err = listener.ActionGetLogs(*conn, message)
			return
		}
	}

	failureMessage := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_FILES,
		Destination: message.Source,
		Message:     "response:false;operation:unknown;message:Operation " + message.Message + " not supported, others errors: ",
	}

	if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Operation " + message.Message + " not supported")); err != nil {
		fmt.Println(err)
		failureMessage.Message += err.Error()
		if err = kernelHandler.WriteServer(*conn, &failureMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason "+err.Error())); err != nil {
				fmt.Println(err)
			}
		}
	}

	return
}



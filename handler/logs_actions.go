package handler

import (
	"fmt"
	"strings"

	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
	"github.com/andresbelo12/fileOS/model"
)

func (listener ClientListener) ActionGetLogs(connection *kernelModel.ClientConnection, message *kernelModel.Message) (err error) {
	messageBody := strings.Split(message.Message, ":")

	failureMessage := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_FILES,
		Destination: message.Source,
		Message:     "response:false;operation:log;message:Log file for module " + messageBody[1] + " not found because ",
	}

	if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager attempting to read file for module: "+messageBody[1])); err != nil {
		fmt.Println(err)
		failureMessage.Message += err.Error()
		if err = kernelHandler.WriteServer(connection, &failureMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
				return
			}
		}
		return
	}

	if logString, err := listener.LogSystem.ReadLogFile(messageBody[1]); err == nil{
		if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager succesfully read file: "+messageBody[1])); err != nil {
			fmt.Println(err)
		}
		responseMessage := kernelModel.Message{
			Command:     kernelModel.CMD_SEND,
			Source:      kernelModel.MD_FILES,
			Destination: message.Source,
			Message:     "response:true;operation:log;message:"+logString,
		}
		fmt.Println(responseMessage)
		if err = kernelHandler.WriteServer(connection, &responseMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
			}
		}
	}else{
		failureMessage.Message += err.Error()
		if err = kernelHandler.WriteServer(connection, &failureMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
			}
		}
		
		if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager failed at read file: "+message.Source+" reason ")); err != nil {
			fmt.Println(err)
		}
	}

	return
}

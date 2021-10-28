package handler

import (
	"fmt"
	"strings"

	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
	"github.com/andresbelo12/fileOS/model"
)

func (listener ClientListener) ActionCreateFolder(connection *kernelModel.ClientConnection, message *kernelModel.Message) (err error) {
	messageBody := strings.Split(message.Message, ":")

	failureMessage := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_FILES,
		Destination: message.Source,
		Message:     "response:false;operation:create;message:Directory " + messageBody[1] + " not created reason ",
	}

	if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager attempting to create folder: "+messageBody[1])); err != nil {
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

	if operationResponse := listener.LogSystem.CreateFolder(messageBody[1]); operationResponse.Success {

		if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager succesfully created folder: "+messageBody[1])); err != nil {
			fmt.Println(err)
		}

		responseMessage := kernelModel.Message{
			Command:     kernelModel.CMD_SEND,
			Source:      kernelModel.MD_FILES,
			Destination: message.Source,
			Message:     "response:true;operation:create;message:Directory " + messageBody[1] + " created",
		}
		if err = kernelHandler.WriteServer(connection, &responseMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
			}
		}
	} else {
		failureMessage.Message += operationResponse.Message
		if err = kernelHandler.WriteServer(connection, &failureMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
			}
		}

		if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager failed at create folder: "+messageBody[1]+" reason "+operationResponse.Message)); err != nil {
			fmt.Println(err)
		}
	}

	return
}

func (listener ClientListener) ActionDeleteFolder(connection *kernelModel.ClientConnection, message *kernelModel.Message) (err error) {
	messageBody := strings.Split(message.Message, ":")

	failureMessage := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_FILES,
		Destination: message.Source,
		Message:     "response:false;operation:delete;message:Directory " + messageBody[1] + " not deleted reason ",
	}

	if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager attempting to delete folder "+messageBody[1])); err != nil {
		failureMessage.Message += err.Error()
		if err = kernelHandler.WriteServer(connection, &failureMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
				return
			}
		}
		return
	}

	if operationResponse := listener.LogSystem.DeleteFolder(messageBody[1]); operationResponse.Success {

		if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager succesfully deleted folder: "+messageBody[1])); err != nil {
			fmt.Println(err)
		}

		responseMessage := kernelModel.Message{
			Command:     kernelModel.CMD_SEND,
			Source:      kernelModel.MD_FILES,
			Destination: message.Source,
			Message:     "response:true;operation:delete;message:Directory " + messageBody[1] + " deleted",
		}
		if err = kernelHandler.WriteServer(connection, &responseMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
			}
		}
	} else {
		failureMessage.Message += operationResponse.Message
		if err = kernelHandler.WriteServer(connection, &failureMessage); err != nil {
			if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("Response not sent to "+message.Source+" reason: "+err.Error())); err != nil {
				fmt.Println(err)
			}
		}

		if err = listener.LogSystem.WriteLog(message.Destination, model.StringToLog("File manager failed at delete folder: "+messageBody[1]+" reason: "+operationResponse.Message)); err != nil {
			fmt.Println(err)
		}
	}

	return
}

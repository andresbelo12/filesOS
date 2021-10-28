package model

import (
	"time"

	"github.com/andresbelo12/KernelOS/model"
)

const(
	timeLayout = "02-Jan-2006 15:04:05"
)

type (
	Log struct {
		Datetime time.Time
		Message string
	}

	OperationResponse struct {
		Success bool
		Message string
	}
)

func OperationResponseToLog(operationResponse *OperationResponse)(log Log){
	log.Datetime = time.Now()
	log.Message = operationResponse.Message
	return
}

func StringToLog(message string)(log Log){
	log.Datetime = time.Now()
	log.Message = message
	return
}

func MessageToLog(message *model.Message)(log Log){
	log.Datetime = time.Now()
	log.Message = "Module " + message.Source + " sent message of type " + message.Command + " to module: " + message.Destination + " message: " + message.Message
	return
}

func (log Log) ToString()(logString string){
	logString = log.Datetime.Format(timeLayout) + " -> " + log.Message + "\n"
	return
}
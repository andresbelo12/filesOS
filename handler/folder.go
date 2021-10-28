package handler

import (
	"os"

	"github.com/andresbelo12/fileOS/model"
)

const (
	DefaultWorkspacePath = "/jalopezb/logs"
	PERMISSIONS          = 0700
)

func (logSystem *LogSystem) CreateFolder(folderName string, firstLog model.Log) *model.OperationResponse {
	folderPath := DefaultWorkspacePath + "/" + folderName
	response := model.OperationResponse{}

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, PERMISSIONS); err != nil {
			response.Message = "Create folder " + folderName + " failed because: " + err.Error()
			response.Success = false
		} else {
			response.Message = "Create folder " + folderName + " successful"
			response.Success = true
			logSystem.WorkspacePath = folderPath
			logSystem.CreateLogFiles(firstLog, "create")
		}
	} else {
		response.Message = "Create folder " + folderName + " failed because already exists "
		response.Success = false
	}

	return &response
}

func (logSystem *LogSystem) DeleteFolder(folderName string, firstLog model.Log) *model.OperationResponse {
	folderPath := DefaultWorkspacePath + "/" + folderName
	response := model.OperationResponse{}

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		response.Message = "Create folder " + folderName + " failed because folder does not exist"
		response.Success = false
	} else {
		if err := os.RemoveAll(folderPath); err != nil {
			response.Message = "Create folder " + folderName + " failed because: " + err.Error()
			response.Success = false
		} else {
			response.Message = "Create folder " + folderName + " successful"
			response.Success = true
			logSystem.WorkspacePath = DefaultWorkspacePath
		}
	}

	return &response
}

func (logSystem LogSystem) CreateDefaultWorkspace() *model.OperationResponse {
	response := model.OperationResponse{}

	if _, err := os.Stat(DefaultWorkspacePath); os.IsNotExist(err) {
		if err := os.MkdirAll(DefaultWorkspacePath, PERMISSIONS); err != nil {
			response.Message = "Create default workspace failed because: " + err.Error()
			response.Success = false
		} else {
			response.Message = "Create default workspace successful"
			response.Success = true
		}
	} else {
		response.Message = "Create default workspace failed because already exists"
		response.Success = false
	}

	return &response
}

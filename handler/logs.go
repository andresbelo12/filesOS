package handler

import (
	"os"

	kernelModel "github.com/andresbelo12/KernelOS/model"
	"github.com/andresbelo12/fileOS/model"
)

type LogSystem struct {
	WorkspacePath string
	ModuleFiles   map[string]*os.File
}

func CreateLogSystem() (logSystem LogSystem) {
	logSystem.WorkspacePath = DefaultWorkspacePath
	logSystem.ModuleFiles = make(map[string]*os.File)
	return
}

func (logSystem LogSystem) InitLogSystem() (err error) {
	firstLog := model.OperationResponseToLog(logSystem.CreateDefaultWorkspace())
	err = logSystem.CreateLogFiles(firstLog)
	return

}

func (logSystem LogSystem) CreateLogFiles(firstLog model.Log) (err error) {
	var modules = [3]string{kernelModel.MD_GUI, kernelModel.MD_FILES, kernelModel.MD_KERNEL}

	for _, module := range modules {
		moduleFilePath := logSystem.WorkspacePath + "/" + module + "_LOGS.txt"
		moduleFile, err := os.OpenFile(moduleFilePath, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}

		logSystem.ModuleFiles[module] = moduleFile

		if module == kernelModel.MD_FILES {
			moduleFile.WriteString(firstLog.ToString())
		}

		moduleFile.WriteString(model.StringToLog(module + " file created at default workspace: " + logSystem.WorkspacePath).ToString())

		if module == kernelModel.MD_KERNEL {
			moduleFile.WriteString(model.StringToLog("Kernel is initializing all modules").ToString())
		}

	}
	return
}

func (logSystem LogSystem) WriteLog(fileName string, log model.Log) (err error) {
	_, err = logSystem.ModuleFiles[fileName].WriteString(log.ToString())
	return
}

package handler

import (
	"os"

	"github.com/andresbelo12/fileOS/model"
)

const(
	MD_GUI = "GUI"
	MD_KERNEL = "Kernel"
	MD_FILES = "FileManager"
)

type LogSystem struct {
	WorkspacePath string
	ModuleFiles map[string]*os.File
	
}

func CreateLogSystem()(logSystem LogSystem) { 
	logSystem.WorkspacePath = DefaultWorkspacePath 
	logSystem.ModuleFiles = make(map[string]*os.File)
	return
}

func (logSystem LogSystem) InitLogSystem() (err error) {
	firstLog := model.OperationResponseToLog(logSystem.CreateDefaultWorkspace())
	err = logSystem.CreateLogFiles(firstLog)
	return

}

func (logSystem LogSystem) CreateLogFiles(firstLog model.Log)(err error){
	var modules = [3]string{MD_GUI,MD_FILES,MD_KERNEL}

	for _, module := range(modules){
		moduleFilePath := logSystem.WorkspacePath + "/" + module + ".txt"
		moduleFile, err := os.OpenFile(moduleFilePath, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil{
			return err
		}

		if module == MD_FILES{
			moduleFile.WriteString(firstLog.ToString())
		}

		moduleFile.WriteString(model.CreateLog(module + " file created at default workspace: " + logSystem.WorkspacePath).ToString())

		if module == MD_KERNEL{
			moduleFile.WriteString(model.CreateLog("Kernel is initializing all modules").ToString())
		}
		
		
	}
	return
}

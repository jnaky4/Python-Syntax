package main

import (
	"Go/files"
)

func main() {
	dir := "/Users/Z004X7X/Git/syntax/Go/multithreading/books"

	startDir, err := files.BuildDirectoryStructure(dir)
	if err != nil {
		println(err.Error())
		return
	}

	files.Root = startDir //todo remove

	for {
		selectedDir, selectedFile, err := files.DisplayDirectoryNavigation(startDir)
		if err != nil {
			println(err.Error())
			return
		}

		if selectedDir != nil {
			files.HandleDirOperation(files.DirectorySelect(selectedDir), selectedDir)
		} else {
			files.HandleFileOperation(files.FileSelect(selectedFile), selectedFile)
		}
	}

	//TODO SCP to file, run file and delete itself

}

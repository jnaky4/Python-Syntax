package main

import (
	"Go/files"
	"fmt"
	"os/exec"
	"path/filepath"
)

// RunSSHCommand executes the SSH command to get directory listing
func RunSSHCommand(hostname, username, keyFile, directory string) (string, error) {
	//ignore := "--ignore="
	//ignore := "--ignore=backup --ignore=environments --ignore=nginx --ignore=platform --ignore=stores --ignore=sshkey_backup --ignore=config --ignore=htpasswd"
	cmd := exec.Command("ssh", "-i", keyFile, fmt.Sprintf("%s@%s", username, hostname), "sudo ls -lhaR "+directory)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil

}

func RunLocalCommand(directory string) (string, error) {
	cmd := exec.Command("ls", "-lhaR", directory)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {

	directory := "/home"

	files.Root = &files.Directory{
		Name:           directory,
		Files:          make(map[string]files.FileInfo),
		Subdirectories: make(map[string]*files.Directory),
		Parent: &files.Directory{
			Name:           filepath.Dir(directory),
			Files:          make(map[string]files.FileInfo),
			Subdirectories: make(map[string]*files.Directory),
			Links:          make(map[string]files.Link),
		},
	}

	files.Root.Parent.Subdirectories[directory] = files.Root

	//SSSH(actualHostname, directory)
	SSSH(expectedHostname, directory)

	//command, err := RunLocalCommand(directory)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	////println(command)

	//localDir := files.ParseDirectoryStructure(command)
	//err = files.SaveDirectoryToFile(&localDir, fmt.Sprintf("./LocalDirStructure.txt"))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//files.CompareRecursion(&expectedDirStructure, &actualDirStructure)

}

func SSSH(hostName string, directory string) {
	//Run the SSH command and get the actualOutput

	username := ""
	keyFile := ""

	actualOutput, err := RunSSHCommand(hostName, username, keyFile, directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse the actualOutput and build the directory structure
	//_ = files.ParseDirectoryStructure(actualOutput)
	actualDirStructure := files.ParseDirectoryStructure(actualOutput)

	//fmt.Printf("%+v\n", actualDirStructure)

	fpath := fmt.Sprintf("./%sDirOutput.txt", hostName)
	err = files.SaveDirectoryToFile(&actualDirStructure, fpath)
	if err != nil {
		fmt.Printf("failed to save file %s -> %v\n", fpath)
		return
	}
}

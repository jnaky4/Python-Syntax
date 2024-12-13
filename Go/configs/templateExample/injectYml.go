package main

import (
	"fmt"
	"os"
	"text/template"
)

type PodData struct {
	Name          string
	App           string
	ContainerName string
	Image         string
}

func createPod(data PodData) {

	// Load the template
	tmpl, err := template.ParseFiles("/Users/Z004X7X/Git/syntax/Go/configs/kubernetes/pod.yaml")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}

	//// Execute the template with the data
	//err = tmpl.Execute(os.Stdout, data)
	//if err != nil {
	//	fmt.Println("Error executing template:", err)
	//}

	file, err := os.Create(fmt.Sprintf("/Users/Z004X7X/Git/syntax/Go/configs/kubernetes/%s.yaml", data.App))
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	defer file.Close() // Ensure the file is closed after writing

	// Execute the template with the data and write the result to the file
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Template saved to output.yaml")

}

func main() {
	createPod(PodData{
		Name:          "my-pod",
		App:           "my-app",
		ContainerName: "nginx-container",
		Image:         "nginx:latest",
	})
}

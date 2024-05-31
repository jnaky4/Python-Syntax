package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	//url1 := "http://localhost:8080/sgdm/environments/llama/rs/foo.json"
	//url1 := "http://localhost:8080/home/cloud-user/unimatrix/environments/t0001/imageSpec/foo.json"
	url1 := "http://localhost:8080/sgdm/environments/zebra/rs/truscanyolorealtime3-v015.json"
	// Get the current working directory
	getwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Construct the file path to the JSON file
	filepath := path.Join(getwd, "RandomProjects", "rsNetworkOverhead", "NetworkDistance", "fullrs.json")

	// Read the JSON file and get the payload
	payload, err := readFile(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Make the POST request with basic authentication
	err = postDataWithBasicAuth(url1, " toss_user_key", "toss_pass_key", payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func postDataWithBasicAuth(url, username, password string, payload []byte) error {
	println("Posting")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	fmt.Println("POST request successful")

	return nil
}

func readFile(filepath string) ([]byte, error) {
	// Read the file at the specified filepath
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

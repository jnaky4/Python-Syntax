package main

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
)

func main() {
	endpoint := "localhost:9000"
	accessKeyID := "poke"
	secretAccessKey := "pokemon123"
	useSSL := false
	bucket := "pokemon"
	filePath := "001Bulbasaur.png"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	path := filepath.Join(getwd, filePath)
	println(path)
	err = minioClient.FGetObject(context.Background(), bucket, "pokemon\\"+filePath, path, minio.GetObjectOptions{})
	if err != nil {
		println(err.Error())
	}
}

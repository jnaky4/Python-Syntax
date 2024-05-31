package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
)

//go:embed images
var images embed.FS

var port = ":8090"

func serveFile(w http.ResponseWriter, req *http.Request){
	file := filepath.Base(req.URL.Path)
	_, err := strconv.Atoi(file)
	if err != nil {
		http.Error(w, "NaN", 404)
	} else {
		filePath := path.Join("images", "pokemon", fmt.Sprintf("%s.png", file))
		http.ServeFile(w, req, filePath)
	}
}

func main() {
	////#1 A basic file handler see : https://eli.thegreenplace.net/2022/serving-static-files-and-web-apps-in-go/
	//handler := http.FileServer(http.Dir("images"))
	//log.Fatal(http.ListenAndServe(port, handler))

	////#2 file handler that strips prefix route and replaces with images folder
	//fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("images")))
	//http.Handle("/static/", fileHandler)
	//log.Fatal(http.ListenAndServe(port, nil))


	////#3 embedding the images as part of the binary
	//// We want to serve static content from the root of the 'images' directory,
	//// but go:embed will create a FS where all the paths start with 'public/...'.
	//// Using fs.Sub we "cd" into 'public' and can serve files relative to it.
	//imageFS, err := fs.Sub(images, "images")
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//
	////Instead of serving the local filesystem on the /images route, we instead embed the public directory into our binary and
	////use Go's virtual filesystem adapter to serve it. This server behaves similarly to the previous one, but it does
	////not require the public directory to be next to its binary anymore - this directory is instead embedded into the
	////binary. It still needs to be on disk when the binary is built, of course, but not when it's run. So we've built a
	////complete web application into a single Go binary.
	//fileHandler := http.StripPrefix("/images/", http.FileServer(http.FS(imageFS)))
	//http.Handle("/images/", fileHandler)
	//log.Fatal(http.ListenAndServe(port, nil))


	//#4 serveFile from handler function
	http.HandleFunc("/file/", serveFile)
	log.Fatal(http.ListenAndServe(port, nil))
}


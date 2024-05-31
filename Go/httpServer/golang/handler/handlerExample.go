package main

import (
	h "Go/httpserver/helpers"
	"log"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet{
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed) ,http.StatusMethodNotAllowed)
		return
	}
	if err := h.WriteJson(w, http.StatusOK, h.Envelope{"header": req.Header}); err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func main(){
	//default serveMux uses a global variable, its possible to modify that and inject handlers
	reqMultiplexer := http.NewServeMux() //Creating a locally scoped mux prevents handler routes from being injected
	reqMultiplexer.HandleFunc("/headers", headers)

	log.Fatal(http.ListenAndServe(":8090", reqMultiplexer))
}


package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
	Handles all the "rough" routing
	set routing for entire applications
	when a network request is received, go to correct controller to be processed
 */
//takes the responsibility to create newUserController
func RegisterControllers(){
	uc := newUserController()

	//takes a pattern and a handler variable of type http.Handler
	http.Handle("/users", *uc)
	//anything /users/* will be handled
	http.Handle("/users/", *uc)

}
func encodeResponseAsJSON(data interface{}, w io.Writer){
	enc := json.NewEncoder(w)
	err := enc.Encode(data)
	if err != nil {
		return
	}

}

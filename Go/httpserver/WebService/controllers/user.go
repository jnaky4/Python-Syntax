package controllers

import (
	"awesomeProject/httpserver/WebService/models"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)


type userController struct{
	//handles resource request on the users collection(all users)
	//manipulate individual users based on the URI path
	userIDPattern *regexp.Regexp

}

/*	Passing Object userController Binds the function as a Method of userController class
		returns ResponseWriter, Request


	implements Handler interface: htptps://golang.org/pkg/net/http/#Handler
			has ServeHTTP, params match and are ordered correctly, this method
			auto implements Handler interface
 */
//serve as a traffic cop to decide which method to hand off to
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//working with entire Users Collection
	if r.URL.Path == "/users"{
		switch r.Method{
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w,r)
		//case http.MethodDelete:
		//	uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	//working with individual user
	} else {
		//compares incoming information
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
			case http.MethodGet:
				uc.get(id, w)
			case http.MethodPut:
				uc.put(id, w, r)
			case http.MethodDelete:
				uc.delete(id, w)
			default:
				w.WriteHeader(http.StatusNotFound)
		}

	}
	//can auto convert string auto to byte, byte is just an alias for string
	//w.Write([]byte("Hello from the User Controller!"))

}
func (uc *userController) getAll(w http.ResponseWriter, r *http.Request){
	encodeResponseAsJSON(models.GetUsers(), w)
}
func (uc *userController) get(id int, w http.ResponseWriter){
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u,w)
}
func (uc *userController) post(w http.ResponseWriter, r *http.Request){
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}
func (uc *userController) put(id int, w http.ResponseWriter, r * http.Request){
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User Object"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}
func (uc *userController) delete(id int, w http.ResponseWriter){
	err := models.RemoveUserByID(id)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (uc *userController) parseRequest(r *http.Request) (models.User, error){
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

//typical naming convention for constructor function
//returns pointer to userController
	//prevents unnecessary copy through call by value
func newUserController() *userController {
	//using "address of" variable (&)
		//permissible with struct data types, can immediately take the address of it
	return &userController{
		//this is a local variable
		//regex uses backtick
		//pattern match: /user/number
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
		//normally when we leave scope, we lose track of the memory being allocated by the function
		//go will recognize if we are returning the address of a local variable
		//address is automatically promoted and won't lose address

	}
}
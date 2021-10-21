package middleware

import (
	"encoding/json"
	"fmt"
	models2 "go-postgres/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	statusBadReq = 400
	statusIse    = 500
)

type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Handler struct {
	Repo models2.Repository
}

func sendErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	res := response{
		ID:      0,
		Message: err.Error(),
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Printf("Unable to encode the response body.  %v", err)
	}
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models2.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Printf("Unable to decode the request body.  %v", err)
		sendErrorResponse(w, err, statusBadReq)
		return
	}

	insertID, err := h.Repo.InsertUser(user)

	if err != nil {
		fmt.Printf("Unable to create user.  %v", err)
		sendErrorResponse(w, err, statusIse)
		return
	}

	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Printf("Unable to encode the response body.  %v", err)
	}
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("Unable to convert the string into int.  %v", err)
		sendErrorResponse(w, err, statusBadReq)
		return
	}

	user, err := h.Repo.GetUser(id)

	if err != nil {
		fmt.Printf("Unable to get user. %v", err)
		sendErrorResponse(w, err, statusIse)
		return
	}

	if err = json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Unable to encode the response body.  %v", err)
	}
}

func (h Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {

	users, err := h.Repo.GetAllUsers()

	if err != nil {
		fmt.Printf("Unable to get all users. %v", err)
		sendErrorResponse(w, err, statusIse)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Printf("Unable to encode the response body.  %v", err)
	}

}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("Unable to convert the string into int.  %v", err)
		sendErrorResponse(w, err, statusBadReq)
		return
	}

	var user models2.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Printf("Unable to decode the request body.  %v", err)
		sendErrorResponse(w, err, statusBadReq)
		return
	}

	updatedRows, err := h.Repo.UpdateUser(id, user)

	if err != nil {
		fmt.Printf("Unable to update user. %v", err)
		sendErrorResponse(w, err, statusIse)
		return
	}
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      id,
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Printf("Unable to encode the response body.  %v", err)
	}
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Printf("Unable to convert the string into int.  %v", err)
		sendErrorResponse(w, err, statusBadReq)
		return
	}

	deletedRows, err := h.Repo.DeleteUser(id)
	if err != nil {
		fmt.Printf("Unable to delete user. %v", err)
		sendErrorResponse(w, err, statusIse)
		return
	}

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      id,
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Printf("Unable to encode the response body.  %v", err)
	}
}

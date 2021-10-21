package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go-postgres/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testDb struct{}

func (t testDb) InsertUser(u models.User) (int, error) {
	if u == user {
		return 1, nil
	}
	return 0, errors.New("users are not equal")
}

func (t testDb) GetUser(i int) (models.User, error) {
	if i == id {
		return user, nil
	}
	return models.User{}, errors.New("ids are not equal")
}

func (t testDb) GetAllUsers() ([]models.User, error) {
	return []models.User{user}, nil
}

func (t testDb) UpdateUser(i int, u models.User) (int64, error) {
	if u == user && i == id {
		return 1, nil
	}
	return 0, errors.New("users or ids are not equal")
}

func (t testDb) DeleteUser(i int) (int64, error) {
	if i == id {
		return 1, nil
	}
	return 0, errors.New("ids are not equal")
}

var (
	testHandler = Handler{Repo: testDb{}}
	user        = models.User{
		Name:   "aaa",
		Rating: 99,
	}
	id = 1
)

func TestHandler_CreateUser(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := "/api/user"
	testJson, _ := json.Marshal(user)
	req, err := http.NewRequest("GET", uri, bytes.NewReader(testJson))
	if err != nil {
		t.Error(err)
	}
	testHandler.CreateUser(resp, req)
	if resp.Result().StatusCode != 200 {
		r := response{}
		json.NewDecoder(resp.Body).Decode(&r)
		t.Error(r.Message)
	}
}

func TestHandler_GetUser(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := fmt.Sprintf("/api/user/%v", id)
	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	testHandler.GetUser(resp, req)
	if resp.Result().StatusCode != 200 {
		r := response{}
		json.NewDecoder(resp.Body).Decode(&r)
		t.Fatal(r.Message)
	}
}

func TestHandler_GetAllUser(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := "/api/user"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	testHandler.GetAllUser(resp, req)
	if resp.Result().StatusCode != 200 {
		r := response{}
		json.NewDecoder(resp.Body).Decode(&r)
		t.Fatal(r.Message)
	}
}

func TestHandler_UpdateUser(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := fmt.Sprintf("/api/user/%v", id)
	testJson, _ := json.Marshal(user)
	req, err := http.NewRequest("PUT", uri, bytes.NewReader(testJson))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	testHandler.UpdateUser(resp, req)
	if resp.Result().StatusCode != 200 {
		r := response{}
		json.NewDecoder(resp.Body).Decode(&r)
		t.Fatal(r.Message)
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := fmt.Sprintf("/api/user/%v", id)
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	testHandler.DeleteUser(resp, req)
	if resp.Result().StatusCode != 200 {
		r := response{}
		json.NewDecoder(resp.Body).Decode(&r)
		t.Fatal(r.Message)
	}
}

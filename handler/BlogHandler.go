package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"blogs/model"
	"blogs/service"

	//"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		log.Println("Error while parsing json", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.CreateBlog(&blog)
	if err != nil {
		println("Error while creating a new blog")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
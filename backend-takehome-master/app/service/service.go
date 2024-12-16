package service

import (
	"app/usecase"
	"net/http"
)

type Service struct {
	Usecase usecase.Usecases
}

type Services interface {
	// user
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)

	// blog
	CreateBlogPost(w http.ResponseWriter, r *http.Request)
	GetBlogPost(w http.ResponseWriter, r *http.Request)
	GetAllBlogPost(w http.ResponseWriter, r *http.Request)
	UpdateBlogPost(w http.ResponseWriter, r *http.Request)
	DeleteBlogPost(w http.ResponseWriter, r *http.Request)

	//comments
	CreateComment(w http.ResponseWriter, r *http.Request)
	GetComments(w http.ResponseWriter, r *http.Request)
}

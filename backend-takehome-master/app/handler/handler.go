package handler

import (
	"app/service"
	"net/http"
	"strings"
)

type Handler struct {
	Service service.Services
}

type Handlers interface {
	HandleRequest()
}

func (handler *Handler) HandleRequest() {
	// user
	http.HandleFunc("/register", handler.Service.RegisterUser)
	http.HandleFunc("/login", handler.Service.LoginUser)
	http.HandleFunc("/posts", handler.handlePost)
	http.HandleFunc("/posts/", handler.handlePost)
}

func (handler *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) == 4 {
		if r.Method == http.MethodGet {
			handler.Service.GetComments(w, r)
			return
		} else if r.Method == http.MethodPost {
			handler.Service.CreateComment(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if len(segments) == 3 {
		if r.Method == http.MethodGet {
			handler.Service.GetBlogPost(w, r)
			return
		} else if r.Method == http.MethodPut {
			handler.Service.UpdateBlogPost(w, r)
		} else if r.Method == http.MethodDelete {
			handler.Service.DeleteBlogPost(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		if r.Method == http.MethodGet {
			handler.Service.GetAllBlogPost(w, r)
			return
		} else if r.Method == http.MethodPost {
			handler.Service.CreateBlogPost(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

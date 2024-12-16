package service

import (
	"app/entity"
	"fmt"
	"net/http"
	"strconv"
)

func (service *Service) CreateComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	postID, err := parsePostIDFromPath(w, r)
	if err != nil {
		return
	}

	content := r.FormValue("content")
	userIDStr := r.FormValue("user_id")
	userToken := r.FormValue("user_token")

	if content == "" {
		http.Error(w, "Content Cannot Be Empty", http.StatusBadRequest)
		return
	}
	if userIDStr == "" {
		http.Error(w, "User ID Cannot Be Empty", http.StatusBadRequest)
		return
	}
	if userToken == "" {
		http.Error(w, "User Token Cannot Be Empty", http.StatusBadRequest)
		return
	}

	authorID, _ := strconv.Atoi(userIDStr)

	if userAllowed, errMess := authenticateUser(userToken, authorID); !userAllowed {
		http.Error(w, errMess, http.StatusBadRequest)
		return
	}

	message := service.Usecase.CreateNewComment(entity.Comment{
		Content:  content,
		AuthorID: authorID,
		PostID:   postID,
	})

	fmt.Fprint(w, message)
}

func (service *Service) GetComments(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	postID, err := parsePostIDFromPath(w, r)
	if err != nil {
		return
	}

	userIDStr := r.FormValue("user_id")
	userToken := r.FormValue("user_token")

	if userIDStr == "" {
		http.Error(w, "User ID Cannot Be Empty", http.StatusBadRequest)
		return
	}
	if userToken == "" {
		http.Error(w, "User Token Cannot Be Empty", http.StatusBadRequest)
		return
	}

	userID, _ := strconv.Atoi(userIDStr)

	if userAllowed, errMess := authenticateUser(userToken, userID); !userAllowed {
		http.Error(w, errMess, http.StatusBadRequest)
		return
	}

	message := service.Usecase.GetAllComments(postID)

	fmt.Fprint(w, message)
}

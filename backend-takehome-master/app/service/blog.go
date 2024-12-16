package service

import (
	"app/entity"
	"fmt"
	"net/http"
	"strconv"
)

func (service *Service) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	if isPost := checkIsPost(w, r); !isPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	userIDStr := r.FormValue("user_id")
	userToken := r.FormValue("user_token")

	if title == "" {
		http.Error(w, "Title Cannot Be Empty", http.StatusBadRequest)
		return
	}
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

	message := service.Usecase.CreateNewBlogPost(entity.BlogPost{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	})

	fmt.Fprint(w, message)
}

func (service *Service) GetBlogPost(w http.ResponseWriter, r *http.Request) {
	postID, err := parsePostIDFromPath(w, r)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
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

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Please enter valid User ID", http.StatusBadRequest)
		return
	}

	if userAllowed, errMess := authenticateUser(userToken, userID); !userAllowed {
		http.Error(w, errMess, http.StatusBadRequest)
		return
	}

	message := service.Usecase.GetBlogPost(postID)

	fmt.Fprint(w, message)
}

func (service *Service) GetAllBlogPost(w http.ResponseWriter, r *http.Request) {
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

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Please enter valid User ID", http.StatusBadRequest)
		return
	}

	if userAllowed, errMess := authenticateUser(userToken, userID); !userAllowed {
		http.Error(w, errMess, http.StatusBadRequest)
		return
	}

	message := service.Usecase.GetAllBlogPost()

	fmt.Fprint(w, message)
}

func (service *Service) UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	postID, err := parsePostIDFromPath(w, r)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	userIDStr := r.FormValue("user_id")
	userToken := r.FormValue("user_token")
	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" && content == "" {
		http.Error(w, "Please Insert Title or Content that you want to update.", http.StatusBadRequest)
		return
	}
	if userIDStr == "" {
		http.Error(w, "User ID Cannot Be Empty.", http.StatusBadRequest)
		return
	}
	if userToken == "" {
		http.Error(w, "User Token Cannot Be Empty.", http.StatusBadRequest)
		return
	}

	authorID, _ := strconv.Atoi(userIDStr)

	if userAllowed, errMess := authenticateUser(userToken, authorID); !userAllowed {
		http.Error(w, errMess, http.StatusBadRequest)
		return
	}

	message := service.Usecase.UpdateBlogPost(entity.BlogPost{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
		PostID:   postID,
	})

	fmt.Fprint(w, message)
}

func (service *Service) DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	postID, err := parsePostIDFromPath(w, r)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	userIDStr := r.FormValue("user_id")
	userToken := r.FormValue("user_token")

	if userIDStr == "" {
		http.Error(w, "User ID Cannot Be Empty.", http.StatusBadRequest)
		return
	}
	if userToken == "" {
		http.Error(w, "User Token Cannot Be Empty.", http.StatusBadRequest)
		return
	}

	authorID, _ := strconv.Atoi(userIDStr)

	if userAllowed, errMess := authenticateUser(userToken, authorID); !userAllowed {
		http.Error(w, errMess, http.StatusBadRequest)
		return
	}

	message := service.Usecase.DeleteBlogPost(authorID, postID)

	fmt.Fprint(w, message)
}

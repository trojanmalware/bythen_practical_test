package service

import (
	"app/entity"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (service *Service) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if isPost := checkIsPost(w, r); !isPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if name == "" {
		http.Error(w, "Name Cannot Be Empty", http.StatusBadRequest)
		return
	}
	if email == "" {
		http.Error(w, "Email Cannot Be Empty", http.StatusBadRequest)
		return
	}
	if password == "" {
		http.Error(w, "Password Cannot Be Empty", http.StatusBadRequest)
		return
	}

	message := service.Usecase.RegisterUser(entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	})

	fmt.Fprint(w, message)
}

func (service *Service) LoginUser(w http.ResponseWriter, r *http.Request) {
	if isPost := checkIsPost(w, r); !isPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" {
		http.Error(w, "Email Cannot Be Empty", http.StatusBadRequest)
	}
	if password == "" {
		http.Error(w, "Password Cannot Be Empty", http.StatusBadRequest)
	}

	user, message := service.Usecase.LoginUser(entity.User{
		Email:    email,
		Password: password,
	})

	if message == "" {
		expiry := internalTokenCache[user.UserID].CreatedTime.Add(time.Second * sessionExpire)

		if hasLoggedIn := time.Now().Before(expiry); hasLoggedIn {
			http.Error(w, fmt.Sprintf("You have already logged in, your token is %s", internalTokenCache[user.UserID].Token), http.StatusBadRequest)
			return
		}

		token := uuid.NewString()
		internalTokenCache[user.UserID] = UserLogin{
			Token:       token,
			CreatedTime: time.Now(),
		}
		message = fmt.Sprintf("You have successfully logged in, your userID is: %d your token is :%s\nPlease use both to authenticate yourself.", user.UserID, token)
	}

	fmt.Fprint(w, message)
}

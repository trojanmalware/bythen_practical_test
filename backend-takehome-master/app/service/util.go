package service

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const sessionExpire = 7200 // set session expire 2 hour
type UserLogin struct {
	Token       string
	CreatedTime time.Time
}

var internalTokenCache = map[int]UserLogin{}

func checkIsPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func authenticateUser(userToken string, userID int) (bool, string) {
	expiry := internalTokenCache[userID].CreatedTime.Add(time.Second * sessionExpire)

	if hasLoggedIn := time.Now().Before(expiry); !hasLoggedIn {
		return false, "Your session has expired. Please Re-Login."
	}

	if userAuthValid := userToken == internalTokenCache[userID].Token; !userAuthValid {
		return false, "Your token is invalid. Please Re-Login or create a new user."
	}

	return true, ""
}

func parsePostIDFromPath(w http.ResponseWriter, r *http.Request) (postID int, err error) {
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 {
		http.Error(w, "No Valid Post ID Found", http.StatusBadRequest)
		return 0, errors.New("No Valid Post ID Found")
	}

	postIDStr := segments[2]
	postID, err = strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Please enter valid Post ID", http.StatusBadRequest)
		return 0, err
	}

	if postID == 0 {
		http.Error(w, "Please enter valid Post ID", http.StatusBadRequest)
		return 0, errors.New("Invalid Post ID")
	}

	return postID, nil

}

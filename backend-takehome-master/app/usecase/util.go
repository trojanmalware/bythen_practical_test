package usecase

import (
	"app/entity"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

const (
	blogPostFormat = "Title: %s\nContent: %s\nAuthor ID: %d\nAuthor Name: %s\nPost ID: %d\nCreated At: %s\nLast Updated At: %s\nComments:\n"
	commentFormat  = " Author: %s\n Comment: %s\n Post ID: %d\n Comment ID: %d\n Created At: %s\n\n"
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPassword(password string) (string, error) {
	// Generate a hashed password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, password string) bool {
	// Compare the hashed password with the plain password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func parseBlogPostToString(blogPost entity.BlogPost) string {
	resp := fmt.Sprintf(blogPostFormat, blogPost.Title, blogPost.Content, blogPost.AuthorID, blogPost.AuthorName, blogPost.PostID, blogPost.CreatedTime.String(), blogPost.UpdatedTime.String())
	for _, comment := range blogPost.Comments {
		resp = resp + parseCommentsToString(comment)
	}
	return resp
}

func parseCommentsToString(comment entity.Comment) string {
	return fmt.Sprintf(commentFormat, comment.Author, comment.Content, comment.PostID, comment.CommentID, comment.CreatedTime.String())
}

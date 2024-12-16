package usecase

import (
	"app/entity"
	"app/provider"
)

type Usecase struct {
	Provider provider.Providers
}

type Usecases interface {
	// user
	LoginUser(param entity.User) (entity.User, string)
	RegisterUser(user entity.User) string

	// blogpost
	CreateNewBlogPost(blogPost entity.BlogPost) string
	GetBlogPost(postID int) string
	GetAllBlogPost() string
	UpdateBlogPost(blogPost entity.BlogPost) string
	DeleteBlogPost(userID, postID int) string

	// comments
	CreateNewComment(comment entity.Comment) string
	GetAllComments(postID int) string
}

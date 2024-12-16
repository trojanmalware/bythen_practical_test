package usecase

import (
	"app/entity"
	"fmt"
)

func (uc *Usecase) CreateNewBlogPost(blogPost entity.BlogPost) string {
	blogID, err := uc.Provider.InsertBlogPost(blogPost)
	if err != "" {
		return err
	}

	return fmt.Sprintf("Successfully Posted a Blog Post. Blog Post ID: %d", blogID)
}

func (uc *Usecase) GetBlogPost(postID int) string {
	blogPost, err := uc.Provider.GetBlogPost(postID)
	if err != "" {
		return err
	}

	user, err := uc.Provider.GetUserByID(blogPost.AuthorID)
	if err != "" {
		fmt.Println("Failed To Get Author Data")
	}

	comments, err := uc.Provider.GetAllComments(postID)
	if err != "" {
		fmt.Println("Failed To Get Comments")
	}

	blogPost.AuthorName = user.Name
	blogPost.Comments = comments

	return parseBlogPostToString(blogPost)
}

func (uc *Usecase) GetAllBlogPost() string {
	blogPosts, err := uc.Provider.GetAllBlogPost()
	if err != "" {
		return err
	}
	resp := ""

	for _, blogPost := range blogPosts {
		user, err := uc.Provider.GetUserByID(blogPost.AuthorID)
		if err != "" {
			fmt.Println("Failed To Get Author Data")
		}

		comments, err := uc.Provider.GetAllComments(blogPost.PostID)
		if err != "" {
			fmt.Println("Failed To Get Comments")
		}

		blogPost.AuthorName = user.Name
		blogPost.Comments = comments
		resp = resp + parseBlogPostToString(blogPost) + "\n"

	}

	return resp
}

func (uc *Usecase) UpdateBlogPost(blogPost entity.BlogPost) string {
	savedBlogPost, err := uc.Provider.GetBlogPost(blogPost.PostID)
	if err != "" {
		return err
	}

	if blogPost.AuthorID != savedBlogPost.AuthorID {
		return "You are not the author of this post. You cannot Edit it."
	}

	if blogPost.Content == "" {
		blogPost.Content = savedBlogPost.Content
	}

	if blogPost.Title == "" {
		blogPost.Title = savedBlogPost.Title
	}

	err = uc.Provider.UpdateBlogPost(blogPost)
	if err != "" {
		return err
	}
	return "Successfully Updated Blog Post."
}

func (uc *Usecase) DeleteBlogPost(userID, postID int) string {
	savedBlogPost, err := uc.Provider.GetBlogPost(postID)
	if err != "" {
		return err
	}

	if userID != savedBlogPost.AuthorID {
		return "You are not the author of this post. You cannot Delete it."
	}

	err = uc.Provider.DeleteBlogPost(postID)
	if err != "" {
		return err
	}
	return "Successfully Deleted Blog Post."
}

package usecase

import (
	"app/entity"
	"fmt"
)

func (uc *Usecase) CreateNewComment(comment entity.Comment) string {
	user, err := uc.Provider.GetUserByID(comment.AuthorID)
	if err != "" {
		fmt.Println("Failed To Get Author Data.")
	}
	comment.Author = user.Name
	err = uc.Provider.InsertComment(comment)
	if err != "" {
		fmt.Println(err)
		return "Failed To Post a Comment. Please Check Your Post ID"
	}

	return "Successfully Posted a Comment"
}

func (uc *Usecase) GetAllComments(postID int) string {
	comments, err := uc.Provider.GetAllComments(postID)
	if err != "" {
		fmt.Println(err)
		return fmt.Sprintf("Failed To Get Comments for Post: %d", postID)
	}

	resp := ""
	for _, comment := range comments {
		resp = resp + parseCommentsToString(comment)
	}

	return resp
}

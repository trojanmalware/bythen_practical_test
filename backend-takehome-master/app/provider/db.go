package provider

import (
	"app/entity"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func (provider *Provider) InsertUser(name, email, password string) string {
	stmt, err := provider.DB.Prepare(queryInserNewUser)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, password)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return "This email has been used. Please use a different email."
		}
		fmt.Println(err)
		return "Failed to Register User, please try again."
	}
	return ""
}

func (provider *Provider) GetUserByEmail(email string) (entity.User, string) {
	row := provider.DB.QueryRow(querySelectUserbyEmail, email)

	userInfo := entity.User{}

	err := row.Scan(&userInfo.Name, &userInfo.HashPassword, &userInfo.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return userInfo, "No user found with that email. Please check your email format or register."
		} else {
			fmt.Println(err)
			return userInfo, "Failed Getting User Info, please try again."
		}
	}

	return userInfo, ""
}

func (provider *Provider) GetUserByID(userID int) (entity.User, string) {
	row := provider.DB.QueryRow(querySelectUserByID, userID)

	userInfo := entity.User{}

	err := row.Scan(&userInfo.Name, &userInfo.HashPassword, &userInfo.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return userInfo, "No user found with that id. Please check your userID."
		} else {
			fmt.Println(err)
			return userInfo, "Failed Getting User Info, please try again."
		}
	}

	return userInfo, ""
}

func (provider *Provider) InsertBlogPost(blogData entity.BlogPost) (int, string) {
	stmt, err := provider.DB.Prepare(queryInsertBlogPost)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	fmt.Println(blogData.AuthorID)
	result, err := stmt.Exec(blogData.Title, blogData.Content, blogData.AuthorID)
	if err != nil {
		fmt.Println(err)
		return 0, "Error Creating Blog Post, Please Try Again."
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	return int(lastInsertID), ""
}

func (provider *Provider) GetBlogPost(postID int) (blogPost entity.BlogPost, errMess string) {

	row := provider.DB.QueryRow(querySelectBlogPost, postID)

	var createdAt []byte
	var updatedAt []byte

	err := row.Scan(&blogPost.PostID, &blogPost.Title, &blogPost.Content, &blogPost.AuthorID, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.BlogPost{}, fmt.Sprintf("No blog post found with ID %d", postID)
		}
		fmt.Println(err)
		return entity.BlogPost{}, fmt.Sprintf("error getting BlogPost: %v", err)
	}

	blogPost.CreatedTime = parseBytetoTime(createdAt)
	blogPost.UpdatedTime = parseBytetoTime(updatedAt)

	return blogPost, ""
}

func (provider *Provider) GetAllBlogPost() (blogData []entity.BlogPost, errMess string) {

	rows, err := provider.DB.Query(querySelectAllBlogPost)
	if err != nil {
		return blogData, fmt.Sprintf("error getting BlogPosts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var blogPost entity.BlogPost
		var createdAt []byte
		var updatedAt []byte
		err := rows.Scan(&blogPost.PostID, &blogPost.Title, &blogPost.Content, &blogPost.AuthorID, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Sprintf("error getting BlogPost row: %v", err)
		}

		blogPost.CreatedTime = parseBytetoTime(createdAt)
		blogPost.UpdatedTime = parseBytetoTime(updatedAt)
		blogData = append(blogData, blogPost)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Sprintf("error iterating BlogPost rows: %v", err)
	}

	return blogData, ""
}

func (provider *Provider) UpdateBlogPost(blogPost entity.BlogPost) (errMess string) {
	result, err := provider.DB.Exec(queryUpdateBlogPost, blogPost.Title, blogPost.Content, blogPost.PostID)
	if err != nil {
		return fmt.Sprintf("error updating blogspot: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Sprintf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Sprintf("no blog post found with ID %d", blogPost.PostID)
	}

	return ""
}

func (provider *Provider) DeleteBlogPost(postID int) (errMess string) {

	result, err := provider.DB.Exec(queryDeleteBlogPost, postID)
	if err != nil {
		return fmt.Sprintf("error deleting blogpost: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Sprintf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Sprintf("no blog post found with ID %d", postID)
	}
	return ""
}

func (provider *Provider) InsertComment(commentData entity.Comment) string {
	stmt, err := provider.DB.Prepare(queryInsertComment)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("Error Inserting Comment to DB.%v", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentData.PostID, commentData.Author, commentData.Content)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1452 {
			return fmt.Sprintf("Post ID:%d Not Found", commentData.PostID)
		}
		fmt.Println(err)
		return fmt.Sprintf("Error Inserting Comment to DB.%v", err.Error())
	}
	return ""
}

func (provider *Provider) GetAllComments(postID int) (comments []entity.Comment, errMess string) {

	rows, err := provider.DB.Query(querySelectAllComment, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []entity.Comment{}, fmt.Sprintf("No Comments found for Blog Post: %d", postID)
		}
		return comments, fmt.Sprintf("error getting Comments: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment

		var createdAt []byte
		err := rows.Scan(&comment.CommentID, &comment.Author, &comment.Content, &createdAt)
		if err != nil {
			return nil, fmt.Sprintf("error getting Comments row: %v", err)
		}
		comment.CreatedTime = parseBytetoTime(createdAt)
		comment.PostID = postID
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Sprintf("error iterating Comments rows: %v", err)
	}

	return comments, ""
}

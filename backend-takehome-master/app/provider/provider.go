package provider

import (
	"app/entity"
	"database/sql"
	"fmt"
)

var listOfIndex = []index{
	{
		table:  "BlogPost",
		column: "author_id",
		index:  "idx_blogpost_author",
	},
	{
		table:  "BlogPost",
		column: "id",
		index:  "idx_blogpost_id",
	},
	{
		table:  "Comment",
		column: "post_id",
		index:  "idx_comment_post_id",
	},
	{
		table:  "User",
		column: "email",
		index:  "idx_user_email",
	},
	{
		table:  "User",
		column: "id",
		index:  "idx_user_id",
	},
}

type Provider struct {
	DB *sql.DB
}

type Providers interface {
	// user session
	InsertUser(name, email, password string) string
	GetUserByEmail(email string) (entity.User, string)
	GetUserByID(userID int) (entity.User, string)

	//blog post
	InsertBlogPost(blogData entity.BlogPost) (int, string)
	GetBlogPost(postID int) (blogPost entity.BlogPost, errMess string)
	GetAllBlogPost() (blogData []entity.BlogPost, errMess string)
	UpdateBlogPost(blogPost entity.BlogPost) (errMess string)
	DeleteBlogPost(postID int) (errMess string)

	//comments
	InsertComment(commentData entity.Comment) string
	GetAllComments(postID int) (comments []entity.Comment, errMess string)
}

type index struct {
	table  string
	column string
	index  string
}

func (provider *Provider) InitProvider() {
	provider.createTable()
	provider.initIndex()
}

func (provider *Provider) createTable() {
	_, err := provider.DB.Exec(queryCreateUserTable)
	if err != nil {
		fmt.Println(err)
	}

	_, err = provider.DB.Exec(queryCreateBlogspotTable)
	if err != nil {
		fmt.Println(err)
	}

	_, err = provider.DB.Exec(queryCreateCommentsTable)
	if err != nil {
		fmt.Println(err)
	}
}

func (provider *Provider) initIndex() {
	for _, index := range listOfIndex {
		if indexExist := provider.checkIndexExist(index.index, index.table); !indexExist {
			err := provider.createIndex(index.table, index.index, index.column)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (provider *Provider) createIndex(tableName, indexName, columnName string) error {
	// Construct the SQL statement for creating the index with custom parameters
	query := fmt.Sprintf(queryCreateIndex, indexName, tableName, columnName)

	// Execute the query
	_, err := provider.DB.Exec(query)
	if err != nil {
		return err
	}
	fmt.Printf("Index '%s' created successfully on table '%s'.\n", indexName, tableName)
	return nil
}

func (provider *Provider) checkIndexExist(indexName, tableName string) bool {
	databaseName := "maindb"

	var result int
	err := provider.DB.QueryRow(queryCheckIndex, databaseName, tableName, indexName).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}

	// Check if the index exists
	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}

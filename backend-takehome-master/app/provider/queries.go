package provider

const (
	queryCreateUserTable = `CREATE TABLE IF NOT EXISTS User (
	id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)`

	queryCreateBlogspotTable = `CREATE TABLE IF NOT EXISTS BlogPost (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES User(id) ON DELETE CASCADE)`

	queryCreateCommentsTable = `CREATE TABLE IF NOT EXISTS Comment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    post_id INT NOT NULL,
    author_name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES BlogPost(id) ON DELETE CASCADE)`

	queryCreateIndex = `CREATE INDEX %s ON %s (%s)`

	queryInserNewUser      = `INSERT INTO User (name, email, password_hash) VALUES (?, ?, ?)`
	querySelectUserbyEmail = `SELECT name, password_hash,id FROM User WHERE email = ?`
	querySelectUserByID    = `SELECT name, password_hash,email FROM User WHERE id = ?`

	queryCheckIndex = `SELECT 1
		FROM information_schema.statistics
		WHERE table_schema = ? 
		  AND table_name = ? 
		  AND index_name = ?
	`

	queryInsertBlogPost = `INSERT INTO BlogPost (title, content, author_id) VALUES (?, ?, ?)`

	querySelectBlogPost = `SELECT id, title, content, author_id, created_at, updated_at FROM BlogPost WHERE id = ?`

	querySelectAllBlogPost = `SELECT id, title, content, author_id, created_at, updated_at FROM BlogPost`

	queryUpdateBlogPost = `UPDATE BlogPost SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`

	queryDeleteBlogPost = `DELETE FROM BlogPost WHERE id = ?`

	queryInsertComment = `INSERT INTO Comment (post_id, author_name, content) VALUES (?, ?, ?)`

	querySelectAllComment = `SELECT id, author_name, content, created_at FROM Comment WHERE post_id = ?`
)

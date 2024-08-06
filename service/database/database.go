/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		
		// Activate the foreign key
		sqlStmt := `PRAGMA foreign_key=ON`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// Create table User
		sqlStmt = `CREATE TABLE users (
			user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE CHECK(length(username) > 2 AND length(username) < 17),
			user_name TEXT NOT NULL,
			user_surname TEXT NOT NULL
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// Create table Post
		sqlStmt = `CREATE TABLE posts(
			post_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			image BLOB NOT NULL,
			timestamp TEXT NOT NULL,
			FOREIGN KEY(user_id) 
				REFERENCES users(user_id)
					ON DELETE CASCADE
		);` 
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// Create table Comment
		sqlStmt = `CREATE TABLE comments(
			comm_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
			timestamp TEXT NOT NULL,
			caption TEXT NOT NULL,
			FOREIGN KEY(user_id) 
				REFERENCES users(user_id)
					ON DELETE CASCADE
			FOREIGN KEY(post_id) 
				REFERENCES posts(post_id)
					ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// Create table Like
		sqlStmt = `CREATE TABLE likes(
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
			PRIMARY KEY(user_id, post_id),
			FOREIGN KEY(user_id)
				REFERENCES users(user_id)
					ON DELETE CASCADE
			FOREIGN KEY(post_id)
				REFERENCES post(post_id)
					ON DELETE CASCADE
		);` 
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// Create table Follows
		sqlStmt = `CREATE TABLE follows(
			user_id INTEGER NOT NULL, 
			followed_id INTEGER NOT NULL,
			PRIMARY KEY(user_id, followed_id),
			FOREIGN KEY(user_id)
				REFERENCES users(user_id)
					ON DELETE CASCADE
			FOREIGN KEY(followed_id)
				REFERENCES users(user_id)
					ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// Create table Bans
		sqlStmt = `CREATE TABLE bans(
			user_id INTEGER NOT NULL,
			banned_id INTEGER NOT NULL,
			PRIMARY KEY(user_id, banned_id),
			FOREIGN KEY(user_id)
				REFERENCES users(user_id)
					ON DELETE CASCADE
			FOREIGN KEY(banned_id)
				REFERENCES users(user_id)
					ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

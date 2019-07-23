package types

import "github.com/go-sql-driver/mysql"

// Configuration holds the app configuration file
type Configuration struct {
	SSLStatus     string
	Privkey       string
	Fullchain     string
	DBDSN         string
	HTTPSPort     string
	HTTPPort      string
	MigrationsDir string
	DBDSNTest     string
}

// LifeEvent Blog Life Event structure
type LifeEvent struct {
	ShownOrder  string
	StartDate   string
	EndDate     string
	Name        string
	Summary     string
	Description string
}

//LifeEvents Stores serveral Life Events
type LifeEvents []LifeEvent

//BlogPost Stores One Blog Post
type BlogPost struct {
	PostID      int            `json:"postID"`
	PublishedOn mysql.NullTime `json:"publishedOn"`
	PostTitle   string         `json:"postTitle"`
	PostContent string         `json:"postContent"`
}

//BlogPosts Stores serveral Blog Post
type BlogPosts []BlogPost

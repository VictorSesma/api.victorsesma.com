package types

import "github.com/go-sql-driver/mysql"

// Configuration holds the app configuration file
type Configuration struct {
	APISecret string `json:"ApiSecret"`
	Ssl       struct {
		Status    string `json:"Status"`
		Privkey   string `json:"Privkey"`
		Fullchain string `json:"Fullchain"`
		HTTPSPort string `json:"HTTPSPort"`
	} `json:"Ssl"`
	Database struct {
		Cassandra struct {
			UserName      string `json:"UserName"`
			Secret        string `json:"Secret"`
			ServerAddress string `json:"ServerAddress"`
			Namespace     string `json:"Namespace"`
		} `json:"Cassandra"`
		MySQL struct {
			UserName      string `json:"UserName"`
			Password      string `json:"Password"`
			ServerAddress string `json:"ServerAddress"`
			DbName        string `json:"DbName"`
		}
	} `json:"Database"`
	HTTPPort string `json:"HTTPPort"`
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

//BlogPost Stores One Blog Post
type BlogPost struct {
	PostID      int            `json:"postID"`
	PublishedOn mysql.NullTime `json:"publishedOn"`
	PostTitle   string         `json:"postTitle"`
	PostContent string         `json:"postContent"`
}

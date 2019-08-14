package helpers

import (
	"database/sql"
	"log"

	"github.com/leviatan89/api.victorsesma.com/types"
)

// GetBlogPosts will get a list of all blogposts in the DB
func GetBlogPosts(db *sql.DB) (*types.BlogPosts, error) {

	query := `
		SELECT postID, publishedOn, postTitle, postContent
		FROM blog
		WHERE publishedOn IS NOT NULL
		ORDER BY publishedOn DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	result := types.BlogPosts{}

	for rows.Next() {
		post := types.BlogPost{}
		err := rows.Scan(&post.PostID, &post.PublishedOn, &post.PostTitle, &post.PostContent)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &result, nil
}

package services

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/leviatan89/api.victorsesma.com/errors"
	"github.com/leviatan89/api.victorsesma.com/helpers"
	"github.com/leviatan89/api.victorsesma.com/types"
)

// Blog is a handler for all blog services
type Blog struct {
	DB *sql.DB
}

// GetAll will fetch all the blogs posts
func (b *Blog) GetAll(r *http.Request, args *string, reply *types.BlogPosts) error {

	response, err := helpers.GetBlogPosts(b.DB)
	if err != nil {
		log.Println(errors.ErrBlogDB)
		return errors.ErrBlogDB
	}

	*reply = *response
	return nil
}

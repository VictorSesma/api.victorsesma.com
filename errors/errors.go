package errors

import "errors"

var (
	// ErrBlogDB error when there is a problem with the DB
	ErrBlogDB = errors.New("Can't fetch the blog posts")
)

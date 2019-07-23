package helpers

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/leviatan89/api.victorsesma.com/types"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.NoError(t, err)
	}
	defer db.Close()

	validTime, err := time.Parse("2006-01-02 15:04:05", "2019-08-08 08:08:08")
	assert.NoError(t, err)

	rows := sqlmock.NewRows([]string{"postID", "publishedOn", "postTitle", "postContent"}).
		AddRow(1, mysql.NullTime{Time: validTime, Valid: true}, "Super Title", "Super Content").
		AddRow(2, mysql.NullTime{Time: validTime, Valid: false}, "Super Title 2", "Super Content 2")

	mock.ExpectQuery("^SELECT (.+) FROM blog (.+)$").WillReturnRows(rows)

	result, err := GetBlogPosts(db)
	assert.NoError(t, err)

	// mysql time when nil
	falseTime, err := time.Parse("2006-01-02 15:04:05", "0001-01-01 00:00:00")
	assert.NoError(t, err)

	expected := types.BlogPosts{
		types.BlogPost{
			PostID: 1,
			PublishedOn: mysql.NullTime{
				Time:  validTime,
				Valid: true,
			},
			PostTitle:   "Super Title",
			PostContent: "Super Content",
		},
		types.BlogPost{
			PostID: 2,
			PublishedOn: mysql.NullTime{
				Time:  falseTime,
				Valid: false,
			},
			PostTitle:   "Super Title 2",
			PostContent: "Super Content 2",
		},
	}
	assert.Equal(t, expected, *result)
}

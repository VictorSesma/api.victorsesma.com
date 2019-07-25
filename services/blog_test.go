package services_test

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/leviatan89/api.victorsesma.com/services"
	"github.com/leviatan89/api.victorsesma.com/types"
	"github.com/stretchr/testify/assert"
)

func startServers(serverRunning chan bool) error {
	// Mocking DB
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println("error:", err)
		return err
	}
	defer db.Close()

	validTime, err := time.Parse("2006-01-02 15:04:05", "2019-08-08 08:08:08")
	if err != nil {
		log.Println("error:", err)
		return err
	}

	rows := sqlmock.NewRows([]string{"postID", "publishedOn", "postTitle", "postContent"}).
		AddRow(1, mysql.NullTime{Time: validTime, Valid: true}, "Super Title", "Super Content")

	mock.ExpectQuery("^SELECT (.+) FROM blog (.+)$").WillReturnRows(rows)

	// Mocking RPC Server
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	b := new(services.Blog)
	b.DB = db
	s.RegisterService(b, "")
	http.Handle("/rpc", s)
	// time.Sleep(10 * time.Second)
	serverRunning <- true
	http.ListenAndServe(":8080", nil)

	return nil
}

func TestGetAll(t *testing.T) {
	serverRunning := make(chan bool)
	go startServers(serverRunning)
	if channel := <-serverRunning; channel == true {
		log.Println("Json RPC test server running ...")
	}

	params := string(`{"jsonrpc":"2.0","id":123,"method":"Blog.GetAll"}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/rpc", strings.NewReader(params))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	assert.NoError(t, err)

	reply := new(types.BlogPosts)
	err = json2.DecodeClientResponse(res.Body, &reply)
	if err != nil {
		log.Println(err)
		assert.Error(t, errors.New("There is a problem in the remote server. Wrong json returned"))
	}

	validTime, err := time.Parse("2006-01-02 15:04:05", "2019-08-08 08:08:08")
	if err != nil {
		log.Println("error:", err)
		assert.NoError(t, err)
	}

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
	}

	assert.Equal(t, expected, *reply)
}

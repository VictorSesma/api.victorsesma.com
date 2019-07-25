package services_test

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/leviatan89/api.victorsesma.com/services"
	"github.com/leviatan89/api.victorsesma.com/types"
	"github.com/stretchr/testify/assert"
)

func startServersCV(serverCVRunning chan bool) error {
	// Mocking DB
	dbcv, mockcv, err := sqlmock.New()
	if err != nil {
		log.Println("error:", err)
		return err
	}
	defer dbcv.Close()

	rows := sqlmock.NewRows([]string{"description", "end_date", "name", "show_order", "start_date", "summary"}).
		AddRow("Volunteer in AEGEE", "2019-08-08 08:08:08", "Aegee coordinator", "1", "2019-08-08 08:08:08", "summary cool")
	mockcv.ExpectQuery("^SELECT (.+) FROM curriculum_vitae (.+)$").WillReturnRows(rows)

	// Mocking RPC Server
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	cv := new(services.CurriculumVitae)
	cv.DB = dbcv
	s.RegisterService(cv, "")
	http.Handle("/rpcCV", s)
	// time.Sleep(10 * time.Second)
	serverCVRunning <- true

	http.ListenAndServe(":8081", nil)

	return nil
}

func TestGetAllCV(t *testing.T) {

	serverCVRunning := make(chan bool)
	go startServersCV(serverCVRunning)
	if channelCV := <-serverCVRunning; channelCV == true {
		log.Println("Json RPC test server running this ...")
	}

	params := string(`{"jsonrpc":"2.0","id":123,"method":"CurriculumVitae.GetAll"}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8081/rpcCV", strings.NewReader(params))

	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	assert.NoError(t, err)

	reply := new(types.LifeEvents)
	err = json2.DecodeClientResponse(res.Body, &reply)
	if err != nil {
		log.Println(err)
		assert.Error(t, errors.New("There is a problem in the remote server. Wrong json returned"))
	}

	expected := types.LifeEvents{
		types.LifeEvent{
			Description: "Volunteer in AEGEE",
			EndDate:     "2019-08-08 08:08:08",
			Name:        "Aegee coordinator",
			ShownOrder:  "1",
			StartDate:   "2019-08-08 08:08:08",
			Summary:     "summary cool",
		},
	}

	assert.Equal(t, expected, *reply)
}

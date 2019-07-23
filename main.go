package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"

	"github.com/leviatan89/api.victorsesma.com/database"
	"github.com/leviatan89/api.victorsesma.com/services"
)

func indeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The API is up and running. Visit https://victorsesma.com/ from the browser.")
}

func main() {

	log.Println("Connecting to DB and aplying migrations...")

	db := database.ConnectDB()
	defer db.Close()

	blog := new(services.Blog)
	blog.DB = db

	cv := new(services.CurriculumVitae)
	cv.DB = db

	log.Println("Starting RPC server...")
	s := rpc.NewServer()
	// We are using json V2 in this API. This allows, for example, named params. "params:{"a": 1, "b": 2}" vs "params:[1, 2]
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterService(blog, "")
	s.RegisterService(cv, "")
	http.Handle("/rpc", s)
	http.ListenAndServe(":80", nil)

}

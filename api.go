package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func blogpost() string {
	return "blogpost :)"
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

func indeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}
func cvHandler(w http.ResponseWriter, r *http.Request) {
	LifeEvents := make(map[string]LifeEvent)
	LifeEvents["0"] = LifeEvent{"0", "2016-06-05", "Current Job", "Full Stack Developer in SmarterClick.com", "PHP, JavaScript, CSS, HTML", "Full Stack Developer in SmarterClick.com. Building all the back-end and front-end internal systems."}
	LifeEvents["1"] = LifeEvent{"1", "2015-05-05", "2016-06-01", "Intern in WatchFit.com", "IT Intern in WatchFit.com", "I did task of project management and developer, assisting to the CTO."}
	js, err := json.Marshal(LifeEvents)
	// fmt.Println(aLifeEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
	// fmt.Fprintf(w, ) prints to the browser
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yeah")
}

func main() {
	http.HandleFunc("/", indeHandler)
	http.HandleFunc("/cv/", cvHandler)
	http.HandleFunc("/blog/", blogHandler)
	http.ListenAndServe(":8000", nil)
}

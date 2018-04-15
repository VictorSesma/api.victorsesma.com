package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Configuration holds the app configuration file
type Configuration struct {
	APISecret string `json:"ApiSecret"`
	Ssl       struct {
		Status    string `json:"Status"`
		Privkey   string `json:"Privkey"`
		Fullchain string `json:"Fullchain"`
		HTTPSPort string `json:"HTTPSPort"`
	} `json:"Ssl"`
	HTTPPort string `json:"HTTPPort"`
}

func configuration() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Println(configuration)
	// Adding semicolons
	configuration.HTTPPort = ":" + configuration.HTTPPort
	configuration.Ssl.HTTPSPort = ":" + configuration.Ssl.HTTPSPort
	return configuration
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

func blogpost() string {
	return "blogpost :)"
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

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	Configuration := configuration()
	target := "https://"
	if Configuration.Ssl.HTTPSPort != "443" {
		target = target + strings.Replace(req.Host, Configuration.HTTPPort, Configuration.Ssl.HTTPSPort, -1)
	}
	target = target + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", indeHandler)
	http.HandleFunc("/cv/", cvHandler)
	http.HandleFunc("/blog/", blogHandler)
	Configuration := configuration()
	go http.ListenAndServe(Configuration.HTTPPort, http.HandlerFunc(redirect))
	err := http.ListenAndServeTLS(Configuration.Ssl.HTTPSPort, Configuration.Ssl.Fullchain, Configuration.Ssl.Privkey, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

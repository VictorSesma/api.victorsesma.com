package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocql/gocql"
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
	Database struct {
		Cassandra struct {
			UserName      string `json:"UserName"`
			Secret        string `json:"Secret"`
			ServerAddress string `json:"ServerAddress"`
			Namespace     string `json:"Namespace"`
		} `json:"Cassandra"`
	} `json:"Database"`
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
	fmt.Fprintf(w, "The API is up and running. Visit https://victorsesma.com/ from the browser.")
}
func cvHandler(w http.ResponseWriter, r *http.Request) {
	Configuration := configuration()
	cluster := gocql.NewCluster(Configuration.Database.Cassandra.ServerAddress)
	cluster.Keyspace = Configuration.Database.Cassandra.Namespace
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: Configuration.Database.Cassandra.UserName,
		Password: Configuration.Database.Cassandra.Secret,
	}
	session, _ := cluster.CreateSession()
	defer session.Close()
	iter := session.Query(`SELECT description, end_date, name, show_order, start_date, summary FROM api_victorsesma.curriculum_vitae;`).Iter()
	var name, summary, description, showOrder string
	var startDate, endDate time.Time
	var counter = 1
	LifeEvents := make(map[string]LifeEvent)
	for iter.Scan(&description, &endDate, &name, &showOrder, &startDate, &summary) {
		//fmt.Println(w, name)
		startDateFormated := startDate.Format("01/2006")
		endDateFormated := endDate.Format("01/2006")
		LifeEvents[strconv.Itoa(counter)] = LifeEvent{showOrder, startDateFormated, endDateFormated, name, summary, description}
		counter++
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
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

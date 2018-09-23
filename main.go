package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/leviatan89/api.victorsesma.com/types"
)

// var sessionCassandra, errCassandra = connectToCassandra()

func configuration() types.Configuration {
	confPath := os.Getenv("GOPATH") + "/src/github.com/leviatan89/api.victorsesma.com/conf.json"
	file, _ := os.Open(confPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := types.Configuration{}
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

// func connectToCassandra() (*gocql.Session, error) {
// 	Configuration := configuration()
// 	cluster := gocql.NewCluster(Configuration.Database.Cassandra.ServerAddress)
// 	cluster.Keyspace = Configuration.Database.Cassandra.Namespace
// 	cluster.Consistency = gocql.Quorum
// 	cluster.Authenticator = gocql.PasswordAuthenticator{
// 		Username: Configuration.Database.Cassandra.UserName,
// 		Password: Configuration.Database.Cassandra.Secret,
// 	}
// 	session, err := cluster.CreateSession()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return session, err
// }

func blogpost() string {
	return "blogpost :)"
}

func indeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The API is up and running. Visit https://victorsesma.com/ from the browser.")
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

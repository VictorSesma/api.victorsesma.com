package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/leviatan89/api.victorsesma.com/types"
)

func blogHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	configuration := configuration()
	conection := configuration.Database.MySQL.UserName + ":" + configuration.Database.MySQL.Password + "@" + configuration.Database.MySQL.ServerAddress + "/" + configuration.Database.MySQL.DbName
	log.Println("Netork is: ", conection)
	db, err := sql.Open("mysql", conection)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Prepare statement for reading data
	query := `
		SELECT postID, publishedOn, postTitle, postContent
		FROM blog
		WHERE publishedOn IS NOT NULL
		ORDER BY publishedOn DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	var postID int
	var publishedOn mysql.NullTime
	var postTitle, postContent string

	var counter = 1
	BlogPosts := make(map[string]types.BlogPost)
	for rows.Next() {
		err := rows.Scan(&postID, &publishedOn, &postTitle, &postContent)
		if err != nil {
			log.Println(err)
			return
		}
		BlogPosts[strconv.Itoa(counter)] = types.BlogPost{postID, publishedOn, postTitle, postContent}
		counter++
	}
	fmt.Printf("blogHanler DB query iterations took %s", time.Since(start))
	js, err := json.Marshal(BlogPosts)
	if err != nil {
		log.Println("------Error!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
	fmt.Printf("cvHandler took %s", time.Since(start))
}

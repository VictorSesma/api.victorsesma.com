package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leviatan89/api.victorsesma.com/types"
)

func cvHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	configuration := configuration()
	conection := configuration.Database.MySQL.UserName + ":" + configuration.Database.MySQL.Password + "@" + configuration.Database.MySQL.ServerAddress + "/" + configuration.Database.MySQL.DbName
	log.Println("Netork is: ", conection)
	db, err := sql.Open("mysql", conection)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	log.Println("Log1")
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Prepare statement for reading data
	query := `
		SELECT description, end_date, name, show_order, start_date, summary
		FROM curriculum_vitae
		WHERE section_type = 'work_experience'
		ORDER BY show_order;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	var name, summary, description, showOrder string
	var startDate, endDate []uint8
	var counter = 1
	LifeEvents := make(map[string]types.LifeEvent)
	for rows.Next() {
		err := rows.Scan(&description, &endDate, &name, &showOrder, &startDate, &summary)
		if err != nil {
			log.Println(err)
			return
		}
		LifeEvents[strconv.Itoa(counter)] = types.LifeEvent{showOrder, string(startDate), string(endDate), name, summary, description}
		counter++
	}
	fmt.Printf("cvHandler DB query iterations took %s", time.Since(start))
	js, err := json.Marshal(LifeEvents)
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
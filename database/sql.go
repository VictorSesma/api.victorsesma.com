package database

import (
	"database/sql"
	"log"

	"github.com/DavidHuie/gomigrate"
	"github.com/leviatan89/api.victorsesma.com/environment"
)

// ConnectDB will connect to the DB and run the migrations if needed
func ConnectDB() *sql.DB {
	conn, err := sql.Open("mysql", environment.GetCofniguration().DBDSN)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	migrator, err := gomigrate.NewMigrator(conn, gomigrate.Mysql{}, environment.GetCofniguration().MigrationsDir)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	err = migrator.Migrate()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return conn
}

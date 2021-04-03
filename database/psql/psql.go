package psql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/mhd53/quanta-fitness-server/config"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
	DbConn   *sql.DB
)

func ConnectDB() (*sql.DB, error) {
	// Load config file + set vars.
	// TODO: Fix this abs path!
	confs := config.LoadConfg("/home/mo/Desktop/Github/quanta-fitness-server/")
	host = confs.Database.DBHost
	port = confs.Database.DBPort
	user = confs.Database.DBUser
	password = confs.Database.DBPassword
	dbname = confs.Database.DBTest
	fmt.Printf("Loaded Config: host=%s, post=%s, user=%s, pwd=%s, dbname=%s", host, port, user, password, dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("DB URL: " + psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error %s: couldn't connect to db", err)
		return nil, err

	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error %s: couldn't ping db", err)
		return nil, err

	}

	fmt.Printf("Successfully connected to %s!\n", dbname)
	return db, nil
}

package adapters

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/mhd53/quanta-fitness-server/config"
	db "github.com/mhd53/quanta-fitness-server/internal/db/sqlc"
)

var (
	testStore *db.Store
	testDB    *sql.DB
)

func TestMain(m *testing.M) {
	testDB, err := connectDB()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testStore = db.NewStore(testDB)
	os.Exit(m.Run())
}

func connectDB() (*sql.DB, error) {
	// Load config file + set vars.
	confs := config.LoadConfg("../../../")
	host := confs.Database.DBHost
	port := confs.Database.DBPort
	user := confs.Database.DBUser
	password := confs.Database.DBPassword
	dbname := confs.Database.DBTest
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

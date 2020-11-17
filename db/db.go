package db

import(
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "root"
	dbname = "hotel"
)

var (
		db *sql.DB
)

func ConnectToDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	d, err := sql.Open("postgres", psqlInfo)

	if err != nil {
			panic(err)
		}

	err = d.Ping()
	if err != nil {
			panic(err)
		}

	fmt.Println("Successfully connected! to database")
	db = d
}


func GetRooms() *sql.Rows {
	rows, err := db.Query("select id, room_count from houses;")

	if err != nil {
		panic(err)
	}
	return rows
}


func GetRoom(roomId string) *sql.Row {
	row := db.QueryRow("select id, room_count from houses where id=$1", roomId)

	return row
}

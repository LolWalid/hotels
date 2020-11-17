package api

import(
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/mux"

	"lolwalid/server/db"
)

func Listen(port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", roomIndex)
	r.HandleFunc("/rooms", roomIndex)
	r.HandleFunc("/rooms/{id:[0-9]+}", roomShow)
	http.Handle("/", r)
	// http.HandleFunc("/", roomIndex)

	return http.ListenAndServe(port, nil)
}

func roomIndex(w http.ResponseWriter, r *http.Request) {
	rows := db.GetRooms()

	var sb strings.Builder

	for rows.Next() {
		var id int
		var roomCount int
		rows.Scan(&id, &roomCount)
		result := fmt.Sprintf("roomCount : %d for room %d\n",  roomCount, id)
		sb.WriteString(result)
	}
	fmt.Fprintf(w, sb.String())
}

func roomShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	roomId := vars["id"]

	row := db.GetRoom(roomId)
	var sb strings.Builder

	var id int
	var roomCount int
	row.Scan(&id, &roomCount)
	result := fmt.Sprintf("roomCount : %d for room %d\n",  roomCount, id)
	sb.WriteString(result)

	fmt.Fprintf(w, sb.String())
}

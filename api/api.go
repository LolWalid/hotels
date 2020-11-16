package api

import(
    "fmt"
    "net/http"
    "strings"

    "lolwalid/server/db"
)

func RoomIndex(w http.ResponseWriter, r *http.Request) {
    rows := db.GetRooms()

    var sb strings.Builder

    for rows.Next()  {
        var id int
        var roomCount int
        rows.Scan(&id, &roomCount)
        result := fmt.Sprintf("roomCount : %d for room %d\n",  roomCount, id)
        sb.WriteString(result)
    }
    fmt.Fprintf(w, sb.String())
}

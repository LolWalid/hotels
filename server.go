package main

import(
    "fmt"
    "net/http"
    "strings"
)

func RoomIndex(w http.ResponseWriter, r *http.Request) {
    rows := GetRooms()

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

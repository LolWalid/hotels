package main

import(
    "log"
    "net/http"
    "lolwalid/server/db"
    "lolwalid/server/api"
)


func init() {
    defer db.ConnectToDb()
}


func main() {
    http.HandleFunc("/rooms", api.RoomIndex)
    http.HandleFunc("/", api.RoomIndex)

    // http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    // })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

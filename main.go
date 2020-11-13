package main

import(
    "log"
    "net/http"
)


func init() {
    defer connectToDb()
}


func main() {
    http.HandleFunc("/rooms", RoomIndex)
    http.HandleFunc("/", RoomIndex)

    // http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    // })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

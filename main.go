package main

import(
	"log"
	"lolwalid/server/db"
	"lolwalid/server/api"
)


func main() {
	db.ConnectToDb()
	log.Fatal(api.Listen(":8080"))
}

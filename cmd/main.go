package main

import (
	"log"

	"github.com/radimzitka/zitodo-mongo/internal/db"
	"github.com/radimzitka/zitodo-mongo/internal/router"
)

func main() {
	err := db.Connect("mongodb+srv://radimzitka:kQ0hReE0QVHh6aOw@cluster0.4kcnciw.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to database!")

	router.Init()
}

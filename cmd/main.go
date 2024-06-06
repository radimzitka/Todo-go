package main

import (
	"log"
	"os"

	"github.com/radimzitka/zitodo-mongo/internal/app"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"github.com/radimzitka/zitodo-mongo/internal/router"
)

func main() {
	err := app.Init(os.Getenv("CONFIG"))
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Connect(app.State.Cfg.Db.ConnectionString)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to database!")

	router.Init()
}

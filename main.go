package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var app application

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.snip)
	mux.HandleFunc("/process", app.process)
	mux.HandleFunc("/delete", app.delete)
	mux.HandleFunc("/create", app.create)

	app.db.openDB()
	defer app.db.closeDB()

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe("127.0.0.1:3000", mux)
	log.Fatal(err)

}

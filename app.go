package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type application struct {
	db DB
}

func (app *application) snip(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"snippet.html",
	}

	data := app.db.GetById(id)

	ts, _ := template.ParseFiles(files...)
	ts.Execute(w, data)

}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"home_page.html",
	}

	data := templateData{}
	data.Snippets = app.db.GetAll()

	ts, _ := template.ParseFiles(files...)
	ts.Execute(w, data)

}

func (app *application) process(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	r.ParseForm()

	var s Snippet
	s.Title = r.FormValue("textTitle")
	s.Note = r.FormValue("textNote")
	s.ID = id

	app.db.Change(s)

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {

	id := app.db.Create()
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	app.db.DeleteRow(id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

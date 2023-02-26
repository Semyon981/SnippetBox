package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

type templateData struct {
	Snippets []Snippet
}

type Snippet struct {
	ID       int
	Title    string
	Note     string
	DateTime string
}

type DB struct {
	database *sql.DB
}

func (db *DB) GetAll() []Snippet {

	row, err := db.database.Query("SELECT * FROM db.notes ORDER BY id DESC;")
	if err != nil {
		panic(err.Error())
	}

	var res []Snippet

	for row.Next() {
		var s Snippet
		row.Scan(&s.ID, &s.Title, &s.Note, &s.DateTime)
		res = append(res, s)
	}

	return res

}

func (db *DB) GetById(id int) Snippet {

	row, err := db.database.Query("SELECT * FROM db.notes WHERE id = " + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}

	if row.Next() {
		var res Snippet
		row.Scan(&res.ID, &res.Title, &res.Note, &res.DateTime)
		return res
	} else {
		panic("Такой записи нет")
	}

}

func (db *DB) Create() int {
	_, err := db.database.Query("INSERT INTO `db`.`notes` (title, note) VALUES ('Новая запись', '');")
	if err != nil {
		panic(err.Error())
	}

	row, _ := db.database.Query("SELECT id FROM db.notes ORDER BY id DESC LIMIT 1;")
	if row.Next() {
		var id int
		row.Scan(&id)
		return id
	} else {
		panic("Ошибка")
	}

}

func (db *DB) Change(s Snippet) {

	_, err := db.database.Query("UPDATE `db`.`notes` SET `title` = '" + s.Title + "', `note` = '" + s.Note + "' WHERE (`id` = '" + strconv.Itoa(s.ID) + "');")
	if err != nil {
		panic(err.Error())
	}
}

func (db *DB) DeleteRow(id int) {

	_, err := db.database.Query("DELETE FROM db.notes WHERE id = " + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}

}

func (db *DB) openDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "usr", "qwerty123", "localhost", "db")
	db.database, _ = sql.Open("mysql", dsn)
}
func (db *DB) closeDB() {
	db.database.Close()
}

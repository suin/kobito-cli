package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteItemRepository struct {
	dbPath string
}

func NewSqliteItemRepository(dbPath string) (this *SqliteItemRepository) {
	this = new(SqliteItemRepository)
	this.dbPath = dbPath
	return
}

func (this *SqliteItemRepository) Items() (items []*Item, err error) {
	db, err := sql.Open("sqlite3", this.dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("select Z_PK, ZTITLE, ZBODY, ZRAW_BODY from ZITEM where ZIN_TRASH is null order by ZUPDATED_AT desc")

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var html sql.NullString
		var markdown sql.NullString
		if err := rows.Scan(&id, &title, &html, &markdown); err != nil {
			return items, err
		}

		items = append(items, NewItem(id, title, html.String, markdown.String))
	}

	return
}

func (this *SqliteItemRepository) ItemOfId(id int) (item *Item, err error) {
	db, err := sql.Open("sqlite3", this.dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select ZTITLE, ZBODY, ZRAW_BODY from ZITEM where Z_PK = ?")

	if err != nil {
		return
	}

	defer stmt.Close()

	var title string
	var html sql.NullString
	var markdown sql.NullString

	err = stmt.QueryRow(id).Scan(&title, &html, &markdown)

	if err != nil {
		return
	}

	item = NewItem(id, title, html.String, markdown.String)

	return
}

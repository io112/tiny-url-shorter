package db

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/io112/tiny-url-shorter/configs"
	_ "github.com/mattn/go-sqlite3"
)

// URL struct of table url in database.
type URL struct {
	Id   int
	Long string
}

// ErrNotFound is returned by QueryUrl when there is no
// row with given id in db.
var ErrNotFound = errors.New("db: Row not found")

// init creates db and tables if they not exists.
func init() {
	query := "CREATE TABLE IF NOT EXISTS url(" +
		"Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, " +
		"long TEXT)"
	db := getDB()
	defer db.Close()
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

// getDB returns instance of database, must be closed with defer db.Close().
func getDB() *sql.DB {
	db, err := sql.Open("sqlite3", configs.DB_NAME)
	if err != nil {
		panic(err)
	}
	return db
}

// QueryURL returns URL by given id.
func QueryURL(id int64) (*URL, error) {
	db := getDB()
	defer db.Close()
	row := sq.Select("Id", "long").
		From("url").
	Where(sq.Eq{"Id": id}).RunWith(db).QueryRow()

	url := URL{}
	err := row.Scan(&url.Id, &url.Long)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &url, nil
}

// CreateURL creates URL in database and returns id.
func CreateURL(longUrl string) (*int64, error) {
	db := getDB()
	defer db.Close()

	res, err := sq.Insert("url").
		Columns("long").
		Values(longUrl).RunWith(db).Exec()
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

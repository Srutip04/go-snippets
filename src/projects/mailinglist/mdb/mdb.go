package mdb

import(
	"database/sql"
	"log"
	"time"
	"github.com/mattn/go-sqlite3"
)

type EmailEntry struct {
	Id int64
	Email string
	ConfirmedAt *time.Time
	OptOut bool
}

func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE emails (
			id INTEGER PRIMARY KEY,
			email TEXT UNIQUE ,
			confirmed_at INTEGER,
			opt_out INTEGER
		);
	`)
	if err != nil {
		if sqlError, ok := err.(sqlite3.Error); ok {
			//code 1 means table already exists
			if sqlError.Code != 1{
				log.Fatal(err)
			}
		}else{
			log.Fatal(err)
		}
	}
}

func emailEntryFromRow(row *sql.Row) (*EmailEntry, error) {
	var id int64
	var email string
	var confirmedAt int64
	var optOut bool
	err := row.Scan(&id, &email, &confirmedAt, &optOut)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	t := time.Unix(confirmedAt, 0)
	return &EmailEntry{Id: id, Email: email, ConfirmedAt: &t, OptOut: optOut}, nil
}
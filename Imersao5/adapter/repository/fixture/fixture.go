package fixture

import (
	"context"
	"database/sql"
	"io/fs"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/maragudk/migrate"
)

func Up(migrationsDir fs.FS) *sql.DB {

	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	err = migrate.Up(context.Background(), db, migrationsDir)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Down(db *sql.DB, migrationsDir fs.FS) {

	err := migrate.Down(context.Background(), db, migrationsDir)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()

}

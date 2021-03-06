package sqlrow_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jjeffery/sqlrow"
	_ "github.com/mattn/go-sqlite3"
)

// The UserRow struct represents a single row in the users table.
// Note that the sqlrow package becomes more useful when tables
// have many more columns than shown in this example.
type UserRow struct {
	ID         int64 `sql:"primary key autoincrement"`
	GivenName  string
	FamilyName string
}

func Example() {
	db, err := sql.Open("sqlite3", ":memory:")
	exitIfError(err)
	setupSchema(db)

	tx, err := db.Begin()
	exitIfError(err)
	defer tx.Rollback()

	// insert three rows, IDs are automatically generated (1, 2, 3)
	for _, givenName := range []string{"John", "Jane", "Joan"} {
		u := &UserRow{
			GivenName:  givenName,
			FamilyName: "Citizen",
		}
		err = sqlrow.Insert(tx, u, `users`)
		exitIfError(err)
	}

	// get user with ID of 3 and then delete it
	{
		var u UserRow
		_, err = sqlrow.Select(tx, &u, `users`, 3)
		exitIfError(err)

		_, err = sqlrow.Delete(tx, u, `users`)
		exitIfError(err)
	}

	// update family name for user with ID of 2
	{
		var u UserRow
		_, err = sqlrow.Select(tx, &u, `users`, 2)
		exitIfError(err)

		u.FamilyName = "Doe"
		_, err = sqlrow.Update(tx, u, `users`)
		exitIfError(err)
	}

	// select rows from table and print
	{
		var users []*UserRow
		_, err = sqlrow.Select(tx, &users, `
			select {}
			from users
			order by id
			limit ? offset ?`, 100, 0)
		exitIfError(err)
		for _, u := range users {
			fmt.Printf("User %d: %s, %s\n", u.ID, u.FamilyName, u.GivenName)
		}
	}

	// Output:
	// User 1: Citizen, John
	// User 2: Doe, Jane
}

func exitIfError(err error) {
	if err != nil {
		log.Output(2, err.Error())
		os.Exit(1)
	}
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func setupSchema(db *sql.DB) {
	_, err := db.Exec(`
		create table users(
			id integer primary key autoincrement,
			given_name text,
			family_name text
		)
	`)
	exitIfError(err)
}

package database

import (
	"database/sql"
	"fmt"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/structures"
	"log"
	"os"
)
import _ "github.com/mattn/go-sqlite3"

type sqliteStorage struct {
}

func NewSqliteStorage() user.Repository {
	db := sqliteStorage{}
	db.Configure()
	return db
}

func (storage sqliteStorage) Configure() {
	os.Remove("./sqlite.s3db")

	db, err := sql.Open("sqlite3", "file:sqlite.s3db?_auth&_auth_user=admin&_auth_pass=hello&_auth_crypt=ssha512&_auth_salt=this_is_my_salt")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (storage sqliteStorage) Create(user structures.User) (structures.User, error) {
	panic("implement me")
}

func (storage sqliteStorage) IsEmailAvailable(email string) (bool, error) {
	panic("implement me")
}

func (storage sqliteStorage) GetHashedPassword(email string) (string, error) {
	panic("implement me")
}

func (storage sqliteStorage) ChangePassword(email, newPassword string) error {
	panic("implement me")
}

func (storage sqliteStorage) SaveResetPasswordToken(token string) error {
	panic("implement me")
}

func (storage sqliteStorage) VerifyResetPasswordToken(token string) (bool, error) {
	panic("implement me")
}

func (storage sqliteStorage) Find(email string) (structures.User, error) {
	panic("implement me")
}

func (storage sqliteStorage) SaveResetToken(email, resetToken string) error {
	panic("implement me")
}

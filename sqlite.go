package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	sqls := []string{
		"create table foo (id integer not null primary key, name text)",
		"delete from foo",
	}
	for _, sql := range sqls {
		_, err = db.Exec(sql)
		if err != nil {
			fmt.Printf("%q: %s\n", err, sql)
			return
		}
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	rows.Close()

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	rows.Close()

}

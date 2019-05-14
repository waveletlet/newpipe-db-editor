package main

import (
	"fmt"
	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func main() {
	conn, err := sqlite3.Open("newpipe.db")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT * FROM playlists")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	cn := stmt.ColumnNames()
	fmt.Println(cn)

	count := 0
	for {
		row, err := stmt.Step()
		if err != nil {
			fmt.Println(err)
		}
		if !row {
			break
		}
		var uid int
		var name string
		var thumbUrl string
		stmt.Scan(&uid, &name, &thumbUrl)

		fmt.Printf("%v: %v	%v	%v	\n", count, uid, name, thumbUrl)

		count++
	}

}

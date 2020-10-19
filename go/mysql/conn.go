package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "test:test@/test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	rows, err := db.Query("SELECT * FROM users") //
	if err != nil {
		panic(err.Error())
	}

	// columns, err := rows.Columns() // カラム名を取得
	// if err != nil {
	// 	panic(err.Error())
	// }
	// columns = ["1"]
	fmt.Print(rows) //[user_id create_at]

}

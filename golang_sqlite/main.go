package golang_sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into products (model, company, price) values('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

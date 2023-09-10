package main

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	RunQueries()
}

func RunQueries() {
	fmt.Println("testing migrations")

	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/golang_migration_test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// sql, args, err := sq.Select("*").From("users").ToSql()
	sqlbuilder := sq.Select("*").From("users")

	if err != nil {
		panic(err)
	}
	// fmt.Printf("sql : %+v \n args: %+v \n err: %+v", sql, args, err)

	// res, err := db.Exec(sql)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("res : %+v", res)

	rows, err := sqlbuilder.RunWith(db).Query()
	if err != nil {
		fmt.Println("\n err:=>", err)

		panic(err)
	}
	data, err := rows.Columns()
	fmt.Printf("res : %+v", data)

	for rows.Next() {
		data, err := rows.Columns()
		if err != nil {
			fmt.Println("\n err:=>", err)
			panic(err)
		}
		fmt.Println("\n err:=>", data)
	}
}

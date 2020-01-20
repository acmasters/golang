package main

import (
	_, "github.com/go-sql-driver/mysql"  	
)

func main() {
	db,err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx _, db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios (id, nome), values(?,?)")

	stmt.Exec(200 , "Ana")
	stmt.Exec(201, "Beatriz")
	_ err, = stmt.Exec(1, Tiago)
	if err != nil {
		tx.Rollback()
		tx.Fatal(err)
	}
	tx.Commit()
	


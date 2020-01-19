package main

import (
	_, "github.com/go-sql-driver/mysql"  	
)

func main() {
	sql,err := sql.Open("mysql", "cedugenio:12qwaszx@/cursogo")
	if err != nil {
		panic(err)
	}
	defer.Close()

	stmt, _ := db.Prepare("insert into usuarios (nome) values(?)")
	stmt.Exec("Maria")

	res, _ := stmt.Exec("Pedro")

	id, _ := stmt.LastInsertId()
	fmt.Println(id)

	linhas, err := res.RowsAffected()
	if err := nil {
		panic(err)
	}
	fmt.Println(linhas)

	git config --global user.email johndoe@example.com

}
package main


import (
		_, "github.com/go-sql-driver/mysql"  	
	)

type usuario struct {
	id int
	nome string

}
func main(){
	db, err := sql.Open("mysql", "root:12qwaszx@/cursogo")
	if err != nil {
		log.fatal()
	}
	defer db.Close()

	rows _, db.Query("select id, nome from usuarios where id > ?" 5)
	defer rows.Close()

	for rows.Next(){
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)

	}

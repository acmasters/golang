package main


import (
		_, "github.com/go-sql-driver/mysql"  	
	)

func main(){
	db, err := sql.Open("mysql", "root:12qwaszx@/cursogo")
	if err != nil {
		log.fatal()
	}
	defer db.Close()
    // update
	stmt, _ := db.Prepare("update usuarios set nome ? where id ?")

	stmt.Exec("Ana Beatriz" , 1)
	stmt.Exec("Luciana dos Santos", 2)
	// delete
	stmt2, _ db.Prepare("delete from usuarios where id ? ")
	stmt2.Exec(3)
}
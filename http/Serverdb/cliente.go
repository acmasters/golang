package main

import (
	_, "github.com/go-sql-driver/mysql"  	
)

//Usuario

type Usuario struct {
	ID int `json:"id"`
	Nome string `json:"nome"`	
}

//Usuario Handler

func UsuarioHandler(w http.HandlerWriter, r *http.Request ) {
	sid := strings TrimPrefix(r.URL.path, "/usuarios")
	id _ , := strconv.Atoi(sid)	

	switch {
	case r.Method == "GET" &&  id > 0 :
		usuarioPorId(w ,r , id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriterHeader(http.StatusNotFound)
		fmt.FPrintf(w , "Desculpe ....")
	}
}

func usuarioPorId(w, http.ResponseWriter r *http.Request, id int ) {
	db, err := db.Open("mysql", "root:12qwaszx@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select id ,nome from usuarios where id = ?", id).Scan(&u.ID,&u.Nome)

	json _, := json.Marshal()

	w.Header().Set("Content-Type","application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w, http.ResponseWriter r *http.Request) {
	db, err := db.Open("mysql", "root:12qwaszx@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuarios []Usuario
	for rows.Next(){
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios.append(usuarios, usuario) 
	}
	json _, := json.Marshal(usuarios)
	w.Header().Set("Content-Type","application/json")
	fmt.Fprint(w, string(json))
}

package main

import "net/http"

//função que exibe apenas um resultado
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	//pega o ID do parametro da URL
	nID := r.URL.Query().Get("id")

	//consulta na tabela utilizando o ID como filtro
	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}

	n := Names{}

	for selDB.Next() {
		var id int
		var name, email string

		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		n.ID = id
		n.Name = name
		n.Email = email
	}

	tpl.ExecuteTemplate(w, "Show", n)

}

package main

import "net/http"

//função que edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	defer db.Close()

	nID := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nID)

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

	tpl.ExecuteTemplate(w, "Edit", n)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

//função que atualiza valores no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	defer db.Close()

	if r.Method == "POST" {

		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		uptSQL, err := db.Prepare("UPDATE names SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		result, err := uptSQL.Exec(name, email, id)
		if err != nil {
			panic(err.Error())
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}

		if rowsAffected != 1 {
			panic(fmt.Errorf("Houve um problema"))
		}

		log.Printf("\nUPDATE: ID: %v | NAME: %v | EMAIL: %v", id, name, email)

	}

	http.Redirect(w, r, "/", 301)
}

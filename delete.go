package main

import (
	"log"
	"net/http"
)

//função que remove valores no banco de dados
func Delete(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	defer db.Close()

	nID := r.URL.Query().Get("id")

	delSQL, err := db.Prepare("DELETE FROM names WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	delSQL.Exec(nID)

	log.Printf("\nDELETE ID: %v", nID)

	http.Redirect(w, r, "/", 301)
}

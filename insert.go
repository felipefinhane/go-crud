package main

import (
	"log"
	"net/http"
)

//função que insere valores no banco de dados
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	//verifica o METHOD do form passado
	if r.Method == "POST" {

		//pega os campos do formulário
		name := r.FormValue("name")
		email := r.FormValue("email")

		//prepara a SQL e verifica erros
		insForm, err := db.Prepare("INSERT INTO names (name, email) VALUES (?,?)")
		if err != nil {
			panic(err.Error())
		}

		//insere os valores do form no SQL
		result, err := insForm.Exec(name, email)
		if err != nil {
			panic(err.Error())
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}

		//exibe um log com os valores digitados no form
		log.Printf("\nINSERT: ID: %v | Name: %v | Email: %v", id, name, email)
	}

	//redirect para home
	http.Redirect(w, r, "/", 301)
}

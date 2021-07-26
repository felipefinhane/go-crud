package main

import (
	"net/http"
)

//renderizar o arquivo Index
func Index(w http.ResponseWriter, r *http.Request) {

	//abre a conexão com o banco de dados utilizando a função dbConn()
	db := dbConn()
	//fecha a conexão no final da função (defer)
	defer db.Close()

	//realiza a consulta com o banco de dados e trata erros
	selDB, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	//monta a struct para ser utilizada no template
	n := Names{}

	//monta um array para guardar os valores da struct
	res := []Names{}

	//realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {

		//armazena os valores em variáveis
		var id int
		var name, email string

		//faz o Scan do SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		//envia os resultados para a Struct
		n.ID = id
		n.Name = name
		n.Email = email

		//adiciona resultado no array
		res = append(res, n)

	}

	//abre a página index e exibe todos os registrados na tela
	tpl.ExecuteTemplate(w, "Index", res)

}

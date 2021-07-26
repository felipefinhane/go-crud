package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

//variavel que renderiza todos os templates da pasta 'template' independete da extensão
var tpl = template.Must(template.ParseGlob("template/*"))

//scruct utilizada para exibir dados no template
type Names struct {
	ID    int
	Name  string
	Email string
}

//função dbConn, abre a conexão com o banco de dados
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "crud"
	dbHost := "db"
	dbPort := "3306"

	urlConn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, urlConn)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {

	//exibe mensagem que o servidor foi iniciado
	log.Println("server started on: http://localhost:9000")

	//gerencia as URL's
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	//ações
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	//inicia o servidor na porta 9000
	http.ListenAndServe(":9000", nil)
}

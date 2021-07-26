package main

import "net/http"

//função que exibe o formulário para inserir novos dados
func New(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "New", nil)
}

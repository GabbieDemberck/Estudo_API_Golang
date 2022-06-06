package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Na struct as variáveis devem ser nomeadas com maúsculo pra publica, minúsculo pra privada !
type Contato struct {
	ID       int    `json:id`
	Nome     string `json:nome`
	Telefone string `json:telefone`
	Email    string `json:email`
}

var contatos []Contato

func main() {
	rota := mux.NewRouter()

	contatos = append(contatos,
		Contato{ID: 1, Nome: "João", Telefone: "(65)5787-5457", Email: "joaozinplay@hiemail.com"},
		Contato{ID: 2, Nome: "Maria", Telefone: "(65)5787-6565", Email: "mariaplay@hiemail.com"},
		Contato{ID: 3, Nome: "Pedro", Telefone: "(65)5485-5248", Email: "pedroplay@hiemail.com"})
	rota.HandleFunc("/contatos", getContatos).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", rota))
}
func getContatos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contatos)
}

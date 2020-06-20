package main

import (
	"github.com/Vioft/Vault-API"
	"log"
	"net/http"
)

func main() {
	router := Vault_API.NewRouter(Vault_API.AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
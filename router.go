package Vault_API

import (
	"github.com/Vioft/Vault-API/logging"
	"github.com/julienschmidt/httprouter"
)

//Reads from the routes slice to translate the values to httprouter.Handle
func NewRouter(routes Routes) *httprouter.Router {

	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = logging.Logger(handle)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}


package Vault_API

import (
	"github.com/Vioft/Vault-API/handlers/vault-port-scanner"
	"github.com/julienschmidt/httprouter"
)

/*
Define all the routes here.
A new Route entry passed to the routes slice will be automatically
translated to a handler with the NewRouter() function
*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"HealthCheck", "GET", "/healthcheck", vault_port_scanner.HealthCheck},
		Route{"LocalPortScan", "GET", "/scan", vault_port_scanner.ScanLocalHost},
		Route{"DevicePortScan", "POST", "/device_scan/:ip_addr", vault_port_scanner.ScanNetworkDevice},
		Route{"Test", "GET", "/test/:try", vault_port_scanner.TestJson},
	}
	return routes
}
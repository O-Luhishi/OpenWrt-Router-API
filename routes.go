package Vault_API

import (
	"github.com/Vioft/Vault-API/handlers/vault-config"
	"github.com/Vioft/Vault-API/handlers/vault-port-scanner"
	"github.com/Vioft/Vault-API/handlers/vault-network-mapper"
	vault_speed "github.com/Vioft/Vault-API/handlers/vault-speed"
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
		// Port-Scanner Module
		Route{"LocalPortScan", "GET", "/portscan/local", vault_port_scanner.ScanLocalHost},
		Route{"NetworkPortScan", "POST", "/portscan/:ip_addr", vault_port_scanner.ScanNetworkDevice},
		Route{"Test", "GET", "/test/:try", vault_port_scanner.TestJson},

		// Vault-Config Module
		Route{"HealthCheck", "GET", "/healthcheck", vault_config.HealthCheck},
		Route{"GetSystemInfo", "GET", "/config/getsysteminfo", vault_config.GetSystemInfo},
		Route{"GetBoardInfo", "GET", "/config/getboardinfo", vault_config.GetBoardInfo},
		Route{"GetWanStatus", "GET", "/config/getwanstatus", vault_config.GetWanStatus},
		Route{"GetLanStatus", "GET", "/config/getlanstatus", vault_config.GetLanStatus},
		Route{"GetWirelessStatus", "GET", "/config/getwirelessstatus", vault_config.GetWirelessStatus},
		Route{"GetServiceList", "GET", "/config/getservicelist", vault_config.GetServiceList},
		Route{"GetInterfaceList", "GET", "/config/getinterfacelist", vault_config.GetInterfaceList},
		Route{"GetNetworkConfig", "GET", "/config/getnetworkconfig", vault_config.GetNetworkConfig},
		Route{"GetWlanClients", "GET", "/config/getwlanclients", vault_config.GetWlanClients},

		// Vault-Network-Mapper Module
		Route{"GetConnectedClients", "GET", "/networkmap/getconnectedclients", vault_network_mapper.GetConnectedDevices},

		Route{"GetDownloadSpeed", "GET", "/get/downloadspeed", vault_speed.Get_Download_Speed},
	}
	return routes
}

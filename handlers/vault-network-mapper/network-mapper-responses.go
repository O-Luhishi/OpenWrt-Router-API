package vault_network_mapper

import "net/http"

type client struct {
	Hostname string `json:"name"`
	MAC      string `json:"mac_address"`
	IP       string `json:"ip_address"`
	STATUS	 bool	`json:"status"`
}

type connection struct {
	Clients []*client `json:"Clients"`
}

func clientResponse(connectedClients []byte, w http.ResponseWriter) {
	response := connectedClients
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

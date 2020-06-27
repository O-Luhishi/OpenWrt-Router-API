package vault_network_mapper

import "net/http"

type client struct {
	Hostname string `json:"Hostname"`
	IP       string `json:"IP"`
	MAC      string `json:"MAC"`
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

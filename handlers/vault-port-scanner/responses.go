package vault_port_scanner

import "net/http"

func clientResponse(connectedClients []byte, w http.ResponseWriter) {
	response := connectedClients
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}


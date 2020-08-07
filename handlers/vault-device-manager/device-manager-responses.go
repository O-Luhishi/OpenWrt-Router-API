package vault_device_manager

import "net/http"

func banClientResponse(status []byte, w http.ResponseWriter) {
	response := status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

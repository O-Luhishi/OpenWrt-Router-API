package vault_speed

import "net/http"

func speedResponse(speed []byte, w http.ResponseWriter) {
	response := speed
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

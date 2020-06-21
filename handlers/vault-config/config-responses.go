package vault_config

import (
	"net/http"
)

func UbusResponse(s func() ([]byte, error), w http.ResponseWriter) {
	stat, err := s()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte(`{"error":"ubus failed"}`))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(stat)
	}
}

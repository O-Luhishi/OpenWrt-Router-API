package vault_config

import (
	"net/http"
)

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

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

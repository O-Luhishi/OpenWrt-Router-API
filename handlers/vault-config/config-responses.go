package vault_config

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type HealthCheckResponse struct {
	// Reserved field to add some meta information to the API response
	Data interface{} `json:"Status"`
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

func healthcheckResponse(status []byte, w http.ResponseWriter) {
	response := status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Writes the response as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&HealthCheckResponse{Data: m}); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JsonErrorResponse{Error: &ApiError{Status: errorCode, Title: errorMsg}})
}

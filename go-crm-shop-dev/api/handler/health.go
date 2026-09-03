package handler

import (
	"encoding/json"
	"net/http"
)

func Health(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(response).Encode(map[string]string{"status": "ok"})
}

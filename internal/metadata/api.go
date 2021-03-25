package metadata

import (
	"encoding/json"

	"net/http"
)

type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

var healthStauts = HealthCheck{Version: "1.0", Status: "ok"}

func CheckApiHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthStauts)
}

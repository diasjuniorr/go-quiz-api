package metadata

import (
	"encoding/json"

	"net/http"
)

func CheckApiHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthStatus)
}

package metadata

type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

var healthStauts = HealthCheck{Version: "1.0", Status: "ok"}

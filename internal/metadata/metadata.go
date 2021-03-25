package metadata

type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

var healthStatus = HealthCheck{Version: "1.0", Status: "ok"}

package metadata

type HealthCheck struct {
	API     string `json:"api"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

var healthStatus = HealthCheck{API: "quizapp", Version: "1.0", Status: "ok"}

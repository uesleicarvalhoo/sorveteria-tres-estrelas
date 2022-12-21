package healthcheck

type Status string

const (
	StatusUp   Status = "UP"
	StatusDown Status = "DOWN"
)

type HealthStatus struct {
	Version  string `json:"version"`
	Status   Status `json:"status"`
	Database Status `json:"database"`
	Cache    Status `json:"cache"`
}

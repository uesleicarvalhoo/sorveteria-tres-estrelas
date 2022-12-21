package healthcheck

type Status string

const (
	StatusUp   Status = "UP"
	StatusDown Status = "DOWN"
)

type HealthStatus struct {
	App      Status `json:"app"`
	Database Status `json:"database"`
	Cache    Status `json:"cache"`
}

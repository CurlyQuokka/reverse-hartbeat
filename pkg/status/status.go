package status

const (
	Running = "running"
	Stopped = "stopped"
)

type Info struct {
	ID      string `json:"id"`
	Value   string `json:"value"`
	Message string `json:"message,omitempty"`
}

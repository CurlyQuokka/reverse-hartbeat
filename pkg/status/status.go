package status

const (
	Running = "running"
	Stopped = "stopped"
)

type Info struct {
	Value string `json:"value"`
}

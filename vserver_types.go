package hetzner

const (
	VServerCommand_Start    = "start"
	VServerCommand_Stop     = "stop"
	VServerCommand_Shutdown = "shutdown"
)

type VServerCommandRequest struct {
	ServerIP string
	Type     string `url:"type"`
}

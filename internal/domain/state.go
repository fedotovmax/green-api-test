package domain

type InstanceState string

const (
	StateNotAuthorized InstanceState = "notAuthorized"
	StateAuthorized    InstanceState = "authorized"
	StateBlocked       InstanceState = "blocked"
	StateSleepMode     InstanceState = "sleepMode"
	StateStarting      InstanceState = "starting"
	StateYellowCard    InstanceState = "yellowCard"
)

type InstanceStateResponse struct {
	StateInstance InstanceState `json:"stateInstance"`
}

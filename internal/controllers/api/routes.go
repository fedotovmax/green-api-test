package api

const (
	routeSendMessage      = "/send-message"
	routeSendFileByURL    = "/send-file-by-url"
	routeGetSettings      = "/get-settings"
	routeGetInstanceState = "/get-instance-state"
)

const prefix = "/api"

const (
	RouteSendMessage      = prefix + routeSendMessage
	RouteSendFileByURL    = prefix + routeSendFileByURL
	RouteGetSettings      = prefix + routeGetSettings
	RouteGetInstanceState = prefix + routeGetInstanceState
)

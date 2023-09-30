package eventdecoder

import "fmt"

// Service events
const (
	SocketHubEvent = iota
	MediaHubEvent
)

func GetEventTypeName(eventID int32) string {
	// Extract service
	service := eventID >> 24

	var serviceName, domainName, eventName string

	switch service {
	case SocketHubEvent:
		serviceName = "socketHubEvent"
		domainName, eventName = getSocketHubDomainAndEventName(eventID)
	case MediaHubEvent:
		serviceName = "mediaHubEvent"
		domainName, eventName = getMediaHubDomainAndEventName(eventID)
	default:
		return "Unknown Service"
	}

	return fmt.Sprintf("%s.%s.%s", serviceName, domainName, eventName)
}

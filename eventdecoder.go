package eventdecoder

import "fmt"

const (
	roomMessage = iota
	chatMessage
	pluginMessage
	pollMessage
	chatChannelMessage
)

const (
	// ChatMessage events
	getChatMessages = iota
	sendChatMessageToRoom
	sendChatMessageToPeers
	editChatMessage
	deleteChatMessage
	getPaginatedMessages
	sendChatMessageToChannel
	searchChannelMessages
	getAllChatChannels
	markChannelIndexAsRead
)

const (
	// PollMessage events
	createPoll = iota
	getPolls
	votePoll
	pollUpdateBroadcast
)

const (
	// RoomMessage events
	getPeerInfo = iota
	updatePeerInfo
	getRoomPeersInfo
	joinRoom
	leaveRoom
	getRoomInfo
	updateRoomInfo
	closeRoom
	LivestreamingStartedEvent
	LivestreamingStoppedEvent
	LivestreamingErroredEvent
	getStagePeers
	getStageRequests
	requestStageAccess
	cancelStageRequest
	grantStageAccess
	denyStageAccess
	roomCountBroadcast
	joinStage
	leaveStage
	getConnectedRoomsDump
	createConnectedRooms
	deleteConnectedRooms
	movePeers
	transferPeer
	movedPeer
	connectedRoomsUpdated
	connectedRoomsDeleted
	getAllAddedParticipants
	broadcastToRoom
	kick
	kickAll
	transcript
	getWaitingRoomRequests
	acceptWaitingRoomRequests
	waitingRoomRequestAccepted
	denyWaitingRoomRequests
	waitingRoomRequestDenied
	peerStatusUpdate
)

const (
	// PluginMessage events
	getPlugins = iota
	addPlugin
	enablePluginForRoom
	disablePluginForPeers
	enablePluginForPeers
	disablePluginForRoom
	removePlugin
	customPluginEventToRoom
	customPluginEventToPeers
	storeInsertKeys
	storeGetKeys
	storeDeleteKeys
	storeDelete
)

const (
	// ChatChannelMessage events
	createChatChannel = iota
	getChatChannel
	deprecatedGetAllChatChannels
	getChannelMembers
	updateChatChannel
)

func GetEventTypeName(eventID int32) string {
	// Extract domain and specific event parts
	domain := eventID >> 16
	specificEvent := eventID & 0xffff

	var domainName, eventName string

	switch domain {
	case roomMessage:
		domainName = "roomMessage"
		eventName = getRoomMessageEventName(specificEvent)
	case chatMessage:
		domainName = "chatMessage"
		eventName = getChatMessageEventName(specificEvent)
	case pluginMessage:
		domainName = "pluginMessage"
		eventName = getPluginMessageEventName(specificEvent)
	case pollMessage:
		domainName = "pollMessage"
		eventName = getPollMessageEventName(specificEvent)
	case chatChannelMessage:
		domainName = "chatChannelMessage"
		eventName = getChatChannelMessageEventName(specificEvent)
	default:
		return "Unknown event"
	}

	return fmt.Sprintf("Event: %s, Specific Event: %s", domainName, eventName)
}

func getRoomMessageEventName(eventID int32) string {
	switch eventID {
	case getPeerInfo:
		return "getPeerInfo"
	case updatePeerInfo:
		return "updatePeerInfo"
	case getRoomPeersInfo:
		return "getRoomPeersInfo"
	case joinRoom:
		return "joinRoom"
	case leaveRoom:
		return "leaveRoom"
	case getRoomInfo:
		return "getRoomInfo"
	case updateRoomInfo:
		return "updateRoomInfo"
	case closeRoom:
		return "closeRoom"
	case LivestreamingStartedEvent:
		return "LivestreamingStartedEvent"
	case LivestreamingStoppedEvent:
		return "LivestreamingStoppedEvent"
	case LivestreamingErroredEvent:
		return "LivestreamingErroredEvent"
	case getStagePeers:
		return "getStagePeers"
	case getStageRequests:
		return "getStageRequests"
	case requestStageAccess:
		return "requestStageAccess"
	case cancelStageRequest:
		return "cancelStageRequest"
	case grantStageAccess:
		return "grantStageAccess"
	case denyStageAccess:
		return "denyStageAccess"
	case roomCountBroadcast:
		return "roomCountBroadcast"
	case joinStage:
		return "joinStage"
	case leaveStage:
		return "leaveStage"
	case getConnectedRoomsDump:
		return "getConnectedRoomsDump"
	case createConnectedRooms:
		return "createConnectedRooms"
	case deleteConnectedRooms:
		return "deleteConnectedRooms"
	case movePeers:
		return "movePeers"
	case transferPeer:
		return "transferPeer"
	case movedPeer:
		return "movedPeer"
	case connectedRoomsUpdated:
		return "connectedRoomsUpdated"
	case connectedRoomsDeleted:
		return "connectedRoomsDeleted"
	case getAllAddedParticipants:
		return "getAllAddedParticipants"
	case broadcastToRoom:
		return "broadcastToRoom"
	case kick:
		return "kick"
	case kickAll:
		return "kickAll"
	case transcript:
		return "transcript"
	case getWaitingRoomRequests:
		return "getWaitingRoomRequests"
	case acceptWaitingRoomRequests:
		return "acceptWaitingRoomRequests"
	case waitingRoomRequestAccepted:
		return "waitingRoomRequestAccepted"
	case denyWaitingRoomRequests:
		return "denyWaitingRoomRequests"
	case waitingRoomRequestDenied:
		return "waitingRoomRequestDenied"
	case peerStatusUpdate:
		return "peerStatusUpdate"
	default:
		return "Unknown roomMessage event"
	}
}

func getChatMessageEventName(eventID int32) string {
	switch eventID {
	case getChatMessages:
		return "getChatMessages"
	case sendChatMessageToRoom:
		return "sendChatMessageToRoom"
	case sendChatMessageToPeers:
		return "sendChatMessageToPeers"
	case editChatMessage:
		return "editChatMessage"
	case deleteChatMessage:
		return "deleteChatMessage"
	case getPaginatedMessages:
		return "getPaginatedMessages"
	case sendChatMessageToChannel:
		return "sendChatMessageToChannel"
	case searchChannelMessages:
		return "searchChannelMessages"
	case getAllChatChannels:
		return "getAllChatChannels"
	case markChannelIndexAsRead:
		return "markChannelIndexAsRead"
	default:
		return "Unknown chatMessage event"
	}
}

func getPluginMessageEventName(eventID int32) string {
	switch eventID {
	case getPlugins:
		return "getPlugins"
	case addPlugin:
		return "addPlugin"
	case enablePluginForRoom:
		return "enablePluginForRoom"
	case disablePluginForPeers:
		return "disablePluginForPeers"
	case enablePluginForPeers:
		return "enablePluginForPeers"
	case disablePluginForRoom:
		return "disablePluginForRoom"
	case removePlugin:
		return "removePlugin"
	case customPluginEventToRoom:
		return "customPluginEventToRoom"
	case customPluginEventToPeers:
		return "customPluginEventToPeers"
	case storeInsertKeys:
		return "storeInsertKeys"
	case storeGetKeys:
		return "storeGetKeys"
	case storeDeleteKeys:
		return "storeDeleteKeys"
	case storeDelete:
		return "storeDelete"
	default:
		return "Unknown pluginMessage event"
	}
}

func getPollMessageEventName(eventID int32) string {
	switch eventID {
	case createPoll:
		return "createPoll"
	case getPolls:
		return "getPolls"
	case votePoll:
		return "votePoll"
	case pollUpdateBroadcast:
		return "pollUpdateBroadcast"
	default:
		return "Unknown pollMessage event"
	}
}

func getChatChannelMessageEventName(eventID int32) string {
	switch eventID {
	case createChatChannel:
		return "createChatChannel"
	case getChatChannel:
		return "getChatChannel"
	case deprecatedGetAllChatChannels:
		return "deprecatedGetAllChatChannels"
	case getChannelMembers:
		return "getChannelMembers"
	case updateChatChannel:
		return "updateChatChannel"
	default:
		return "Unknown chatChannelMessage event"
	}
}

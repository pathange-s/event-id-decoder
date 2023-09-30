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
	domain := GetDomainEvent(eventID)
	specificEvent := GetDomainSpecificEvent(eventID)

	var domainName, eventName string

	switch domain {
	case roomMessage:
		domainName = "roomMessage"
		eventName = 
getRoomMessageEventName(specificEvent)
	case chatMessage:
		domainName = "chatMessage"
		eventName = 
getChatMessageEventName(specificEvent)
	case pluginMessage:
		domainName = "pluginMessage"
		eventName = 
getPluginMessageEventName(specificEvent)
	case pollMessage:
		domainName = "pollMessage"
		eventName = 
getPollMessageEventName(specificEvent)
	case chatChannelMessage:
		domainName = "chatChannelMessage"
		eventName = 
getChatChannelMessageEventName(specificEvent)
	default:
		return "Unknown event"
	}

	return fmt.Sprintf("Event: %s, Specific Event: %s", 
domainName, eventName)
}

func GetDomainEvent(eventID int32) int32 {
	return eventID >> 16
}

func GetDomainSpecificEvent(eventID int32) int32 {
	return eventID & 0xffff
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


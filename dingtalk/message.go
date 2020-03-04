package dingtalk

type MessageType string

const (
	MessageTypeText       MessageType = "text"
	MessageTypeMarkdown   MessageType = "markdown"
	MessageTypeLink       MessageType = "link"
	MessageTypeActionCard MessageType = "actionCard"
	MessageTypeFeedCard   MessageType = "feedCard"
)

type Message interface {
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

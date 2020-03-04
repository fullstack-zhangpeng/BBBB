package dingtalk

type TextMessage struct {
	MsgType MessageType `json:"msgtype"`
	Text    Text        `json:"text"`
	At      At          `json:"at"`
}

type Text struct {
	Content string `json:"content"`
}

func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		MsgType: MessageTypeText,
		Text:    Text{Content: content},
	}
}

func (m *TextMessage) SetAt(atMobiles []string, isAtAll bool) {
	m.At = At{
		AtMobiles: atMobiles,
		IsAtAll:   isAtAll,
	}
}

package dingtalk

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

type LinkMessage struct {
	MsgType MessageType `json:"msgtype"`
	Link    Link        `json:"link"`
}

func NewLinkMessage(link Link) *LinkMessage {
	return &LinkMessage{
		MsgType: MessageTypeLink,
		Link:    link,
	}
}

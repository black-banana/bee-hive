package hub

type Message struct {
	Author  string      `json:"author,omitempty"`
	Content interface{} `json:"content"`
	Type    string      `json:"type"`
}

func NewMessage() Message {
	return Message{"", "", "msg"}
}

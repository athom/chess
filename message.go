package chess

import "encoding/json"

func NewMessage(data string, myTurn bool) (r *Message) {
	r = &Message{
		UI:     data,
		MyTurn: myTurn,
	}
	return
}

type Message struct {
	UI     string `json:"ui"`
	MyTurn bool   `json:"myTurn"`
}

func (this *Message) ToJson() (r []byte) {
	r, _ = json.Marshal(this)
	return
}

//type IncommingMessage struct {
	//Event   Event
	//Content string
//}

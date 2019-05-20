package hive-pusher-go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/* Message struct */
type reqMessage struct {
	Token string `json:"token"`
}

type resMessage struct {
	Code int `json:"code"`
	Text string `json:"text"`
	Method string `json:"method"`
}

/* Pusher and functions */
type Pusher struct{
	Path string
	Token string
}

func (p Pusher) SetPath(path string) {
	p.Path = path
}

func (p Pusher) SetToken(token string) {
	p.Token = token
}

func (p Pusher) Send() {
	req := reqMessage{
		Token: p.Token,
	}
	buff := new(bytes.Buffer)
	_ = json.NewEncoder(buff).Encode(req)

	resRaw, err := http.Post(p.Path, "application/json; charset=utf-8", buff)

	if err != nil {
		_ = fmt.Errorf("[hive-pusher]%s", err)
	} else {
		var res resMessage
		decoder := json.NewDecoder(resRaw.Body)
		err := decoder.Decode(&res)

		if err != nil {
			_ = fmt.Errorf("[hive-pusher]%s", err)
		} else if res.Code != 200 {
			_ = fmt.Errorf("[hive-pusher]%s", res.Text)
		}
	}
}

package Send2Json

import (
	"bytes"
	"fmt"
	"net/http"
)

type sendJson struct {
	client *http.Client
	url    string
	json   []byte
}

func NewSendJson(client *http.Client, url string, json []byte) *sendJson {
	return &sendJson{
		client: client,
		url:    url,
		json:   json,
	}
}
func (sj *sendJson) Send() error {
	res, err := sj.client.Post(sj.url, "application/json", bytes.NewBuffer(sj.json))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		return fmt.Errorf("http status is not succeed. status is %s\n", res.Status)
	}
	return nil
}

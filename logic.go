package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// send
func (c *Client) send(pushMsg *PushMsg) (*Response, error) {
	sendResponse := &Response{}

	// 参数处理
	jsonByte, err := json.Marshal(pushMsg)
	if err != nil {
		return sendResponse, err
	}

	request, err := http.NewRequest("POST", apiServerUrl, bytes.NewBuffer(jsonByte))
	request.Header.Set("Authorization", fmt.Sprintf("key=%v", client.ApiKey))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return sendResponse, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return sendResponse, err
	}
	sendResponse.StatusCode = response.StatusCode
	if response.StatusCode != 200 {
		return sendResponse, nil
	}

	// 解析响应
	if err := json.Unmarshal([]byte(body), &sendResponse); err != nil {
		return sendResponse, err
	}
	sendResponse.Ok = true
	return sendResponse, nil
}

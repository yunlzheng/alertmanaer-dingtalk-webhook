package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yunlzheng/alertmanaer-dingtalk-webhook/model"
)

// Send send markdown message to dingtalk
func Send(markdown *model.DingTalkMarkdown) (err error) {
	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		"POST",
		"https://oapi.dingtalk.com/robot/send?access_token=ea10c4648177dc44ec540a621a3e431518124880c98c491bc0469ea311242741",
		bytes.NewBuffer(data))

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}

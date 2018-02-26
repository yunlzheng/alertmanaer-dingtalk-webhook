package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yunlzheng/alertmanaer-dingtalk-webhook/model"
	"github.com/yunlzheng/alertmanaer-dingtalk-webhook/transformer"
)

// Send send markdown message to dingtalk
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	var dingTalkRobotURL string

	if robotURL != "" {
		dingTalkRobotURL = robotURL
	} else {
		dingTalkRobotURL = defaultRobot
	}

	req, err := http.NewRequest(
		"POST",
		dingTalkRobotURL,
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

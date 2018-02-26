package transformer

import "github.com/yunlzheng/alertmanaer-dingtalk-webhook/model"

func TransformToMarkdown(notification model.Notification) (err error, markdown *model.DingTalkMarkdown) {

	// groupKey := notification.GroupKey

	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Title: "title",
			Text:  "#### 监控指标\n> 监控描述信息\n\n> ###### 告警时间 \n",
		},
		At: &model.At{
			IsAtAll: false,
		},
	}

	return
}

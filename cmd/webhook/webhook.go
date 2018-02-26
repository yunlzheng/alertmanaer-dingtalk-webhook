package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/yunlzheng/alertmanaer-dingtalk-webhook/model"
	"github.com/yunlzheng/alertmanaer-dingtalk-webhook/notifier"
)

var (
	h            bool
	defaultRobot string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global dingtalk robot webhook, you can overwrite by alert rule with annotations dingtalkRobot")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = notifier.Send(notification, defaultRobot)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})

	})
	router.Run()
}

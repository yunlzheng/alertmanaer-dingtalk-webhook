package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/yunlzheng/alertmanaer-dingtalk-webhook/model"
	"github.com/yunlzheng/alertmanaer-dingtalk-webhook/notifier"
	"github.com/yunlzheng/alertmanaer-dingtalk-webhook/transformer"
)

func main() {
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification
		if err := c.BindJSON(&notification); err == nil {

			err, markdown := transformer.TransformToMarkdown(notification)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			if err = notifier.Send(markdown); err == nil {
				c.JSON(http.StatusOK, gin.H{"message": "notification dates are valid!"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	router.Run()
}

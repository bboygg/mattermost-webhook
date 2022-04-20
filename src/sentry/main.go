package sentry

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pure-juan/mattermost-webhook/src/request"
)

func ReceiveWebhook(c *gin.Context) {
	channel := c.Param("channel")
	var body SentryTriggeredPayload
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}

	var color, text string
	switch body.Level {
	case "error":
		color = "#F44336"
		text = "Okay, Chill out.\nTake a deep breath and let's see\n"
	default:
		color = "#007bff"
		text = "I don't know What is this, Maybe you can figure it out\n\n"
	}

	if len(body.Event.Exception.Values) > 0 {
		text += fmt.Sprintf("```\n%s\n", body.Event.Exception.Values[0].Stacktrace.Frames[0].ContextLine)
		for i := len(body.Event.Exception.Values[0].Stacktrace.Frames) - 1; i > 0; i-- {
			v := body.Event.Exception.Values[0].Stacktrace.Frames[i]
			text += fmt.Sprintf("\tat %s (%s) %d:%d\n", v.Function, v.AbsPath, v.Lineno, v.Colno)
		}
		text += "```"
	}

	var message string
	if body.Message != "" {
		message = body.Message
	} else {
		message = body.Event.Metadata.Value
	}
	request.Sentry(channel, gin.H{
		"attachments": []gin.H{
			{
				"color":      color,
				"title":      "Sentry reported " + body.Level,
				"title_link": body.URL,
				"text":       text,
				"fields": []gin.H{
					{
						"short": true,
						"title": "Project",
						"value": body.ProjectName,
					},
					{
						"short": true,
						"title": "Environment",
						"value": body.Event.Environment,
					},
					{
						"short": true,
						"title": "Type",
						"value": body.Event.Metadata.Type,
					},
					{
						"short": true,
						"title": "Message",
						"value": message,
					},
					{
						"short": true,
						"title": "Platform",
						"value": body.Event.Platform,
					},
					{
						"short": false,
						"title": "Link",
						"value": fmt.Sprintf("[%s](%s)", body.URL, body.URL),
					},
				},
			},
		},
	})
	c.JSON(200, gin.H{})
}

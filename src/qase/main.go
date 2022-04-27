package qase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/pure-juan/mattermost-webhook/src/request"
)

//Run this function in Qase
func ReceiveWebhook(c *gin.Context) {
	channel := c.Param("channel")
	var body BaseQaseTriggeredPayload
	if err := c.BindJSON(&body); err != nil {
		println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	raw, err := json.Marshal(body.Payload)
	if err != nil {
		println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if body.EventName == "run.started" {
		var payload RunTestPayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
			return
		}
		// mapstructure.Decode(body.Payload, &payload)
		request.Qase(channel, gin.H{
			"attachments": []gin.H{
				{
					"title": "Qase Test Run Started",
					"text":  fmt.Sprintf("[%s](https://app.qase.io/run/%s/dashboard/%d)", payload.Title, body.ProjectCode, payload.ID),
					"fields": []gin.H{
						{
							"short": true,
							"title": "Cases_count",
							"value": payload.CasesCount,
						},
						{
							"short": true,
							"title": "Description",
							"value": payload.Description,
						},
						{
							"short": true,
							"title": "Environment",
							"value": payload.Environment,
						},
						// {
						// 	"short": true,
						// 	"title": "Link",
						// 	"value": fmt.Sprintf("[%s](https://app.qase.io/run/%s/dashboard/%d)", body.ProjectCode, body.ProjectCode, payload.ID),
						// },
					},
				},
			},
		})
	} else if body.EventName == "run.completed" {
		var payload CompleteTestPayload
		mapstructure.Decode(body.Payload, &payload)
		request.Qase(channel, gin.H{
			"attachments": []gin.H{
				{
					"title": "Qase Test Run Completed",
					"text":  fmt.Sprintf("[%s](https://app.qase.io/run/%s/dashboard/%d)", "See the Result", body.ProjectCode, payload.ID),
					"fields": []gin.H{
						{
							"short": true,
							"title": "cases",
							"value": payload.Cases,
						},
						{
							"short": true,
							"title": "failed",
							"value": payload.Failed,
						},
						{
							"short": true,
							"title": "passed",
							"value": payload.Passed,
						},
						{
							"short": true,
							"title": "blocked",
							"value": payload.Blocked,
						},
						{
							"short": true,
							"title": "duration",
							"value": payload.Duration / 1000,
						},
						// {
						// 	"short": true,
						// 	"title": "Link",
						// 	"value": fmt.Sprintf("[%s](https://app.qase.io/run/%s/dashboard/%d)", body.ProjectCode, body.ProjectCode, payload.ID),
						// },
					},
				},
			},
		})

	}
	c.JSON(200, gin.H{})
}

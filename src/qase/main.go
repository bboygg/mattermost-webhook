package qase

import (
	"encoding/json"
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
					"title": "Qase Test Run Started ",
					"text":  payload.Title,
					"fields": []gin.H{
						{
							"short": true,
							"title": "cases_count",
							"value": payload.CasesCount,
						},
						{
							"short": true,
							"title": "description",
							"value": payload.Description,
						},
						{
							"short": true,
							"title": "environment",
							"value": payload.Environment,
						},
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
							"title": "duration",
							"value": payload.Duration / 1000,
						},
					},
				},
			},
		})

	}
	c.JSON(200, gin.H{})
}

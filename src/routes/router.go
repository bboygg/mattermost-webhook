package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pure-juan/mattermost-webhook/src/qase"
	"github.com/pure-juan/mattermost-webhook/src/sentry"
)

func Init(r *gin.RouterGroup) {
	InitSentry(r.Group("/sentry"))
	InitQase(r.Group("/qase"))
}

func InitQase(r *gin.RouterGroup) {
	r.POST("/:channel", qase.ReceiveWebhook)
}

func InitSentry(r *gin.RouterGroup) {
	InitSentryMiddleware(r)
	r.POST("/:channel", sentry.ReceiveWebhook)
}

func InitSentryMiddleware(r *gin.RouterGroup) {
}

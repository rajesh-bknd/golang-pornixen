package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/my/repo/src/controllers/channels"
)

func InitChannelRoutes(router *gin.Engine) {
	router.GET(CHANNELS, controller.GetAllChannels)
}

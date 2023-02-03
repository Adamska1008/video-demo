package router

import (
	"github.com/gin-gonic/gin"
	"video_demo/src/controller"
)

func InitRouter(r *gin.Engine) {
	BaseRouter(r)
}

func BaseRouter(r *gin.Engine) {
	r.GET("/douyin/feed", controller.BaseContro.Feed)
	r.POST("/douyin/publish/action", controller.BaseContro.PublishAction)
	r.POST("/douyin/publish/list", controller.BaseContro.PublishList)
	r.Static("/videos", "./res/user_upload_video")
	r.Static("/covers", "./res/user_upload_cover")
}

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"video_demo/src/controller"
)

func InitRouter(r *gin.Engine) {
	BaseRouter(r)
}

func BaseRouter(r *gin.Engine) {
	r.GET("/douyin/feed", controller.BaseContro.Feed)
	r.POST("/douyin/publish/action", controller.BaseContro.PublishAction)
	r.POST("/douyin/publish/list", controller.BaseContro.PublishList)
	r.StaticFS("/videos", http.Dir("./res/user_upload_video"))
	r.StaticFS("/covers", http.Dir("./res/user_upload_cover"))
}

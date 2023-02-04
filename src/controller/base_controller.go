package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"video_demo/src/service"
)

const (
	MaxVideoStreamNumber int = 30
)

var (
	BaseContro = &BaseController{}
)

type BaseController struct{}

func (t *BaseController) Feed(c *gin.Context) {
	var latestTime int64
	if reqTime := c.DefaultQuery("latest_time", ""); reqTime != "" {
		timestamp, _ := strconv.Atoi(reqTime)
		latestTime = int64(timestamp)
	} else {
		latestTime = time.Now().Unix()
	}
	videoList, nextTime, err := service.BasicService.ListVideoBefore(time.Unix(latestTime, 0), MaxVideoStreamNumber)
	data := map[string]interface{}{
		"status_code": 0,
		"status_msg":  "",
		"video_list":  videoList,
		"next_time":   nextTime,
	}
	if err != nil {
		data["status_code"] = 1
	}
	c.AsciiJSON(http.StatusOK, data)
}

func (t *BaseController) PublishAction(c *gin.Context) {
	file, _ := c.FormFile("data")
	title := c.PostForm("title")
	data := map[string]interface{}{
		"status_code": 0,
		"status_msg":  0,
	}
	if err := service.BasicService.SaveVideo(file, title, c); err != nil {
		data["status_code"] = 1
	}
	c.AsciiJSON(http.StatusOK, data)
}

func (t *BaseController) PublishList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	videoList, err := service.BasicService.ListVideoByAuthorId(int64(userId))
	data := map[string]interface{}{
		"status_code": 0,
		"status_msg":  "",
		"video_list":  videoList,
	}
	if err != nil {
		data["status_code"] = 1
	}
	c.AsciiJSON(http.StatusOK, data)
}

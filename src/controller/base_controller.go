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
	var latestTime time.Time
	if reqTime := c.DefaultQuery("latest_time", ""); reqTime != "" {
		timestamp, _ := strconv.Atoi(reqTime)
		latestTime = time.Unix(int64(timestamp), 0)
	} else {
		latestTime = time.Now()
	}
	videoList, nextTime, err := service.BasicService.ListVideoBefore(latestTime, MaxVideoStreamNumber)
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

}

func (t *BaseController) PublishList(c *gin.Context) {

}

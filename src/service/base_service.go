package service

import (
	"time"
	"video_demo/src/model"
)

var (
	BasicService = &BasicServicesImpl{}
)

type RespVideo struct {
}

// 从数据库类型Video转化为Response的Video格式
func (r *RespVideo) fromVideo(video model.Video) {

}

type BasicServices interface {
	ListVideoBefore(time time.Time, limit int) []*RespVideo
}

type BasicServicesImpl struct{}

func (b *BasicServicesImpl) ListVideoBefore(time time.Time, limit int) (videoList []*RespVideo, nextTime int) {
	// todo
}

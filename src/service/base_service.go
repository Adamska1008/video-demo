package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
	"strings"
	"time"
	"video_demo/src/model"
	"video_demo/src/tools"
)

var (
	BasicService = &BasicServicesImpl{}
)

type BasicServices interface {
	ListVideoBefore(time time.Time, limit int) []*RespVideo
}

type BasicServicesImpl struct{}

func (b *BasicServicesImpl) ListVideoBefore(latest time.Time, limit int) (videoList []*RespVideo, nextTime int64, err error) {
	rawVideos, err := model.ListVideoBefore(latest, limit)
	if err != nil {
		return nil, 0, err
	}
	nextTime = time.Now().Unix()
	for _, rawVideo := range rawVideos {
		respVideo, err := fromVideo(rawVideo)
		if err != nil {
			return nil, 0, err
		}
		videoList = append(videoList, respVideo)
		if respVideo.PublishDate < nextTime {
			nextTime = respVideo.PublishDate
		}
	}
	return
}

func (b *BasicServicesImpl) ListVideoByAuthorId(authorId int64) (videoList []*RespVideo, err error) {
	rawVideos, err := model.ListVideoByAuthorId(authorId)
	if err != nil {
		return nil, err
	}
	for _, rawVideo := range rawVideos {
		respVideo, err := fromVideo(rawVideo)
		if err != nil {
			return nil, err
		}
		videoList = append(videoList, respVideo)
	}
	return
}

// SaveVideo demo没有用户部分，所以先统一存到一个目录下
func (b *BasicServicesImpl) SaveVideo(header *multipart.FileHeader, title string, c *gin.Context) error {
	// 实际项目中应当通过token获取用户Id
	var userId int64 = 125794582365478165
	video := model.Video{
		AuthorId:      userId,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		PublishDate:   time.Now(),
	}
	videoId, err := model.AddVideo(&video)
	if err != nil {
		return err
	}
	ext := strings.ToLower(path.Ext(header.Filename))
	dst := fmt.Sprintf("./res/user_upload_video/%v/%v.%v", userId, videoId, ext)
	if err = c.SaveUploadedFile(header, dst); err != nil {
		return err
	}
	if err = tools.ExtractCover(video.AuthorId, video.Id); err != nil {
		return err
	}
	return nil
}

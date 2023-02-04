package model

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
	"video_demo/src/tools"
)

type Video struct {
	Id            int64 `gorm:"primaryKey;autoIncrement"`
	AuthorId      int64
	FavoriteCount int64
	CommentCount  int64
	Title         string
	PublishDate   time.Time
}

func (t *Video) TableName() string {
	return "t_video"
}

// PlayerUrl 地址格式为 "/videos/{author_id}/{video_id}"
func (t *Video) PlayerUrl() string {
	localPath, _ := tools.ObtainVideoPath(t.AuthorId, t.Id)
	ext := path.Ext(localPath)
	return fmt.Sprintf("/videos/%v/%v.%v", t.AuthorId, t.Id, ext)
}

func (t *Video) CoverUrl() string {
	localPath, _ := tools.ObtainCoverPath(t.AuthorId, t.Id)
	if _, err := os.Stat(localPath); err != nil {
		tools.ExtractCover(t.AuthorId, t.Id)
	}
	ext := path.Ext(localPath)
	return fmt.Sprintf("/covers/%v/%v.%v", t.AuthorId, t.Id, ext)
}

func AddVideo(video *Video) (int64, error) {
	if err := db.Create(video).Error; err != nil {
		log.Fatal(err)
		return 0, err
	}
	return video.Id, nil
}

func FindVideoById(videoId int64) (*Video, error) {
	var video Video
	if err := db.First(&video, videoId).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &video, nil
}

func DeleteVideo(videoId int64) error {
	if err := db.Delete(&Video{Id: videoId}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func ListVideoByAuthorId(authorId int64) (videos []*Video, err error) {
	if err = db.Where("author_id = ?", authorId).Find(&videos).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return
}

// ListVideoBefore
// 返回在latest之前发布的视频，最多limit条
// 选择视频的算法未定
func ListVideoBefore(latest time.Time, limit int) ([]*Video, error) {
	var videos []*Video
	if err := db.Where("publish_date < ?", latest).Find(&videos).Error; err != nil {
		return nil, err
	}
	if len(videos) > limit {
		return videos[:limit], nil
	} else {
		return videos, nil
	}
}

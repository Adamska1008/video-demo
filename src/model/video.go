package model

import (
	"fmt"
	"log"
	"time"
)

type Video struct {
	Id            int64 `gorm:"primaryKey;autoIncrement"`
	AuthorId      int64
	FavoriteCount int64
	CommentCount  int64
	Title         string
	PublishDate   *time.Time
}

func (t *Video) TableName() string {
	return "t_video"
}

// PlayerUrl 地址格式为 "/videos/{author_id}/{video_id}"
func (t *Video) PlayerUrl() string {
	return fmt.Sprintf("/videos/%v/%v.mp4", t.AuthorId, t.Id)
}

func (t *Video) CoverUrl() string {
	return fmt.Sprintf("/covers/%v/%v.png", t.AuthorId, t.Id)
}

func AddVideo(video *Video) error {
	if err := db.Create(video).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
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

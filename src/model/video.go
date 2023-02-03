package model

import (
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

func ListVideoBefore(latest time.Time, limit int) []*Video {
	// todo
}

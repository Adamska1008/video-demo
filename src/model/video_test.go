package model

import (
	"fmt"
	"log"
	"testing"
	"time"
	"video_demo/src/config"
)

func buildEnv() error {
	cfg, err := config.Load("../../config.yml")
	if err != nil {
		return err
	}
	InitDB(cfg)
	return nil
}

func TestFindVideoById(t *testing.T) {
	if err := buildEnv(); err != nil {
		t.Error(err)
		return
	}

	var videoId int64 = 597814563214596347
	video, err := FindVideoById(videoId)
	if err != nil {
		t.Error(err)
		return
	}
	if video == nil {
		t.Error("Could not find video")
		return
	}
	fmt.Println(video)
}

func TestAddVideo(t *testing.T) {
	if err := buildEnv(); err != nil {
		t.Error(err)
		return
	}
	current := time.Now()
	video := Video{
		AuthorId:      125794582365478165,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "Test",
		PublishDate:   &current,
	}
	if err := db.Create(&video).Error; err != nil {
		log.Fatal(err)
		return
	}
}

func TestDeleteVideo(t *testing.T) {
	if err := buildEnv(); err != nil {
		t.Error(err)
		return
	}
	if err := DeleteVideo(597814563214596348); err != nil {
		log.Fatal(err)
		return
	}
}

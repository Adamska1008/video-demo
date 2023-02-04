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
	video := Video{
		AuthorId:      125794582365478165,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "Test",
		PublishDate:   time.Now(),
	}
	id, err := AddVideo(&video)
	if err != nil {
		t.Error(err)
		return
	}
	if id == 0 {
		t.Error("Not give correct id")
	}
	fmt.Println(id)
}

func TestDeleteVideo(t *testing.T) {
	if err := buildEnv(); err != nil {
		t.Error(err)
		return
	}
	if err := DeleteVideo(597814563214596349); err != nil {
		log.Fatal(err)
		return
	}
}

func TestListVideoBefore(t *testing.T) {
	if err := buildEnv(); err != nil {
		t.Error(err)
		return
	}
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2023-02-02 16:14:13")
	videos, err := ListVideoBefore(testTime, 30)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(len(videos))
	for _, video := range videos {
		fmt.Println(video)
	}
}

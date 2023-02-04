package tools

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"path/filepath"
)

// ObtainVideoPath
// 通过id获取本地存储的文件路径，通过Glob匹配。
func ObtainVideoPath(authorId, videoId int64) (string, error) {
	noExt := fmt.Sprintf("./res/user_upload_video/%v/%v.*", authorId, videoId)
	files, err := filepath.Glob(noExt)
	if err != nil {
		return "", err
	}
	return files[0], nil
}

// ObtainCoverPath
// 通过id获取本地存储的文件路径，通过Glob匹配。
func ObtainCoverPath(authorId, videoId int64) (string, error) {
	noExt := fmt.Sprintf("./res/user_upload_cover/%v/%v.*", authorId, videoId)
	files, err := filepath.Glob(noExt)
	if err != nil {
		return "", err
	}
	return files[0], nil
}

// ExtractCover
// 通过id获取视频的第一帧，并保存在对应地址下
func ExtractCover(authorId, videoId int64) error {
	buf := bytes.NewBuffer(nil)
	videoPath, err := ObtainVideoPath(authorId, videoId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	coverPath := fmt.Sprintf("./res/user_upload_cover/%v/%v.png", authorId, videoId)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{"gte(n,1)"}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatalf("生成封面失败:\n%v", err)
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatalf("生成封面失败:\n%v", err)
		return err
	}
	err = imaging.Save(img, coverPath)
	if err != nil {
		log.Fatalf("生成封面失败:\n%v", err)
		return err
	}
	return nil
}

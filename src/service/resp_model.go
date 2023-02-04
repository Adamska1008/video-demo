package service

import "video_demo/src/model"

type RespUser struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func fromUser(user *model.User) *RespUser {
	return &RespUser{
		Id:            user.Id,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
}

type RespVideo struct {
	Id             int64     `json:"id"`
	Author         *RespUser `json:"author"`
	PlayUrl        string    `json:"play_url"`
	CoverUrl       string    `json:"cover_url"`
	FavouriteCount int64     `json:"favourite_count"`
	CommentCount   int64     `json:"comment_count"`
	IsFavorite     bool      `json:"is_favorite"`
	Title          string    `json:"title"`
	PublishDate    int64     `json:"publish_date"`
}

// 从数据库类型Video转化为Response的Video格式
func fromVideo(video *model.Video) (*RespVideo, error) {
	user, err := model.FindUserById(video.AuthorId)
	if err != nil {
		return nil, err
	}
	return &RespVideo{
		Id:             video.Id,
		Author:         fromUser(user),
		PlayUrl:        video.PlayerUrl(),
		CoverUrl:       video.CoverUrl(),
		FavouriteCount: video.FavoriteCount,
		CommentCount:   video.CommentCount,
		IsFavorite:     false,
		Title:          video.Title,
		PublishDate:    video.PublishDate.Unix(),
	}, nil
}

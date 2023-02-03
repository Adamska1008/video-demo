package model

import "log"

type User struct {
	Id            int64 `gorm:"primaryKey"`
	Username      string
	FollowCount   int64 // 关注总数
	FollowerCount int64 // 粉丝总数
}

func (*User) TableName() string {
	return "t_user"
}

func FindUserById(id int64) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &user, nil
}

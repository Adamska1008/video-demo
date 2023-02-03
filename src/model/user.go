package model

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Username string
}

func (*User) TableName() string {
	return "t_user"
}

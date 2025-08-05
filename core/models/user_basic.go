package models

type UserBasic struct {
	Id       int64
	Identity string
	Name     string
	Password string
	Email    string
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}

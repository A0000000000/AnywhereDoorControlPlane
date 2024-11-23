package model

type User struct {
	Id       int    `json:"id" gorm:"primary_key" column:"id"`
	Username string `json:"username" column:"username"`
	Password string `json:"password" column:"password"`
}

func (User) TableName() string {
	return "t_user"
}

package model

type Imsdk struct {
	Id            int    `json:"id" gorm:"primary_key" column:"id"`
	UserId        int    `json:"user_id" gorm:"column:user_id" column:"user_id"`
	ImsdkName     string `json:"imsdk_name" gorm:"column:imsdk_name" column:"imsdk_name"`
	ImsdkDescribe string `json:"imsdk_describe" gorm:"column:imsdk_describe" column:"imsdk_describe"`
	ImsdkHost     string `json:"imsdk_host" gorm:"column:imsdk_host" column:"imsdk_host"`
	ImsdkPort     int    `json:"imsdk_port" gorm:"column:imsdk_port" column:"imsdk_port"`
	ImsdkPrefix   string `json:"imsdk_prefix" gorm:"column:imsdk_prefix" column:"imsdk_prefix"`
	ImsdkToken    string `json:"imsdk_token" gorm:"column:imsdk_token" column:"imsdk_token"`
	IsActive      int    `json:"is_active" gorm:"column:is_active" column:"is_active"`
}

func (Imsdk) TableName() string {
	return "t_imsdk"
}

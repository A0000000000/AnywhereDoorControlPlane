package model

type Config struct {
	Id          int    `json:"id" gorm:"primary_key" column:"id"`
	UserId      int    `json:"user_id" gorm:"column:user_id" column:"user_id"`
	TargetId    int    `json:"target_id" gorm:"column:target_id" column:"target_id"`
	Type        int    `json:"type" gorm:"column:type" column:"type"`
	ConfigKey   string `json:"config_key" gorm:"column:config_key" column:"config_key"`
	ConfigValue string `json:"config_value" gorm:"column:config_value" column:"config_value"`
}

func (Config) TableName() string {
	return "t_config"
}

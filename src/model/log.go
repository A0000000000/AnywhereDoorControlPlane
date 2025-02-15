package model

type Log struct {
	Id        int    `json:"id" gorm:"primary_key" column:"id"`
	UserId    int    `json:"user_id" gorm:"column:user_id" column:"user_id"`
	Type      int    `json:"type" gorm:"column:type" column:"type"`
	TargetId  int    `json:"target_id" gorm:"column:target_id" column:"target_id"`
	Timestamp int64  `json:"timestamp" gorm:"column:timestamp" column:"timestamp"`
	Level     int    `json:"level" gorm:"column:level" column:"level"`
	Tag       string `json:"tag" gorm:"column:tag" column:"tag"`
	Log       string `json:"log" gorm:"column:log" column:"log"`
}

func (Log) TableName() string {
	return "t_log"
}

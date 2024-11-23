package model

type Plugin struct {
	Id             int    `json:"id" gorm:"primary_key" column:"id"`
	UserId         int    `json:"user_id" column:"user_id"`
	PluginName     string `json:"plugin_name" column:"plugin_name"`
	PluginDescribe string `json:"plugin_describe" column:"plugin_describe"`
	PluginHost     string `json:"plugin_host" column:"plugin_host"`
	PluginPort     int    `json:"plugin_port" column:"plugin_port"`
	PluginPrefix   string `json:"plugin_prefix" column:"plugin_prefix"`
	PluginToken    string `json:"plugin_token" column:"plugin_token"`
	IsActive       int    `json:"is_active" column:"is_active"`
}

func (Plugin) TableName() string {
	return "t_plugin"
}

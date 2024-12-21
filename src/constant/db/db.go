package db

const (
	EnvDbIp       = "DB_IP"
	EnvDbPort     = "DB_PORT"
	EnvDbUser     = "DB_USER"
	EnvDbPassword = "DB_PASSWORD"
	EnvDbName     = "DB_NAME"

	DefaultDbIp       = "192.168.25.7"
	DefaultDbPort     = "3306"
	DefaultDbUser     = "root"
	DefaultDbPassword = "09251205"
	DefaultDbName     = "anywhere_door"

	DsnTemplate = "%s:%s@tcp(%s:%s)/%s?%s"
	TimeZone    = "time_zone='Asia%2FShanghai'"

	QueryUserIdSQLTemplate = "user_id = ?"
	QueryPluginSQLTemplate = "user_id = ? AND plugin_name = ?"
	QueryImsdkSQLTemplate  = "user_id = ? AND imsdk_name = ?"
)

package db

const (
	EnvDbIp       = "DB_IP"
	EnvDbPort     = "DB_PORT"
	EnvDbUser     = "DB_USER"
	EnvDbPassword = "DB_PASSWORD"
	EnvDbName     = "DB_NAME"

	DefaultDbIp       = "default ip"
	DefaultDbPort     = "default port"
	DefaultDbUser     = "default user"
	DefaultDbPassword = "default password"
	DefaultDbName     = "default db"

	DsnTemplate = "%s:%s@tcp(%s:%s)/%s?%s"
	TimeZone    = "time_zone='Asia%2FShanghai'"

	QueryUserIdSQLTemplate = "user_id = ?"
	QueryPluginSQLTemplate = "user_id = ? AND plugin_name = ?"
	QueryImsdkSQLTemplate  = "user_id = ? AND imsdk_name = ?"
)

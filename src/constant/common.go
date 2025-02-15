package constant

const (
	Username = "username"
	Token    = "token"

	PluginURI         = "/plugin"
	PluginURLTemplate = "http://%s:%d%s/plugin"

	ImsdkURI         = "/imsdk"
	ImsdkURLTemplate = "http://%s:%d%s/imsdk"

	ImsdkConfigURI  = "/imsdk/config"
	PluginConfigURI = "/plugin/config"

	ImsdkLogURL  = "/imsdk/log"
	PluginLogURL = "/plugin/log"

	ContentType     = "content-type"
	ContentTypeJSON = "application/json"

	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"

	TypeControlPlane = 2
	TypePlugin       = 3
	TypeImsdk        = 4

	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4

	HttpStatusSuccess = 200

	GlobalLogUserId = -1
)

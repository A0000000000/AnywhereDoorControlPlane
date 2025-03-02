package apis

import (
	"AnywhereDoorControlPlane/constant"
	"AnywhereDoorControlPlane/constant/code"
	"AnywhereDoorControlPlane/constant/message"
	"AnywhereDoorControlPlane/db"
	"AnywhereDoorControlPlane/model"
	"AnywhereDoorControlPlane/server"
	"github.com/gin-gonic/gin"
)

func InitRegisterServer(logCtx *LogContext, dbCtx *db.DataBaseContext, httpServerCtx *server.HttpServerContext) {
	TAG := "RegisterServer"
	httpServerCtx.Post(constant.ImsdkRegisterURI, func(c *gin.Context) {
		username := c.Request.Header.Get(constant.Username)
		token := c.Request.Header.Get(constant.Token)
		if len(username) == 0 || len(token) == 0 {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.UsernameOrTokenEmpty,
				Msg:  message.UsernameOrTokenEmpty,
				Data: nil,
			})
			return
		}
		user := dbCtx.QueryUser(username)
		if user.Id == -1 {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.UserNotExist,
				Msg:  message.UserNotExist,
				Data: nil,
			})
			return
		}
		var imsdk model.Imsdk
		imsdk.UserId = user.Id
		imsdk.ImsdkToken = token
		var params model.RegisterParams
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.JsonParseError,
				Msg:  message.JsonParseError,
				Data: err.Error(),
			})
			logCtx.E(TAG, "bind json err. url: "+constant.ImsdkRegisterURI+", error: "+err.Error())
			return
		}
		if params.Port <= 0 || params.Port > 65535 || params.Name == "" || params.Host == "" {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.ParamsError,
				Msg:  message.ParamsError,
				Data: nil,
			})
			return
		}
		tempImsdk := dbCtx.QueryImsdk(username, params.Name)
		if tempImsdk.Id != -1 {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NameRepeat,
				Msg:  message.NameRepeat,
				Data: nil,
			})
			return
		}
		imsdk.ImsdkName = params.Name
		imsdk.ImsdkHost = params.Host
		imsdk.ImsdkPort = params.Port
		imsdk.ImsdkPrefix = params.Prefix
		imsdk.IsActive = 1
		id, err := dbCtx.InsertImsdk(imsdk)
		if err != nil {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.RegisterFailed,
				Msg:  message.RegisterFailed,
				Data: err.Error(),
			})
			return
		} else {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.Success,
				Msg:  message.Success,
				Data: id,
			})
		}
	})

	httpServerCtx.Post(constant.PluginRegisterURI, func(c *gin.Context) {
		username := c.Request.Header.Get(constant.Username)
		token := c.Request.Header.Get(constant.Token)
		if len(username) == 0 || len(token) == 0 {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.UsernameOrTokenEmpty,
				Msg:  message.UsernameOrTokenEmpty,
				Data: nil,
			})
			return
		}
		user := dbCtx.QueryUser(username)
		if user.Id == -1 {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.UserNotExist,
				Msg:  message.UserNotExist,
				Data: nil,
			})
			return
		}
		var plugin model.Plugin
		plugin.UserId = user.Id
		plugin.PluginToken = token
		var params model.RegisterParams
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.JsonParseError,
				Msg:  message.JsonParseError,
				Data: err.Error(),
			})
			logCtx.E(TAG, "bind json err. url: "+constant.ImsdkRegisterURI+", error: "+err.Error())
			return
		}
		if params.Port <= 0 || params.Port > 65535 || params.Name == "" || params.Host == "" {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.ParamsError,
				Msg:  message.ParamsError,
				Data: nil,
			})
			return
		}
		tempPlugin := dbCtx.QueryPlugin(username, params.Name)
		if tempPlugin.Id != -1 {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NameRepeat,
				Msg:  message.NameRepeat,
				Data: nil,
			})
			return
		}
		plugin.PluginName = params.Name
		plugin.PluginHost = params.Host
		plugin.PluginPort = params.Port
		plugin.PluginPrefix = params.Prefix
		plugin.IsActive = 1
		id, err := dbCtx.InsertPlugin(plugin)
		if err != nil {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.RegisterFailed,
				Msg:  message.RegisterFailed,
				Data: err.Error(),
			})
			return
		} else {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.Success,
				Msg:  message.Success,
				Data: id,
			})
		}
	})

}

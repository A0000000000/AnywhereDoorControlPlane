package config

import (
	"AnywhereDoorControlPlane/constant"
	"AnywhereDoorControlPlane/constant/code"
	"AnywhereDoorControlPlane/constant/message"
	"AnywhereDoorControlPlane/db"
	"AnywhereDoorControlPlane/model"
	"AnywhereDoorControlPlane/server"
	"github.com/gin-gonic/gin"
)

func InitConfigServer(dbCtx *db.DataBaseContext, httpServerCtx *server.HttpServerContext) {
	httpServerCtx.Post(constant.ImsdkConfigURI, func(c *gin.Context) {
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
		var params model.ConfigParams
		err := c.ShouldBindJSON(&params)
		if err != nil {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.JsonParseError,
				Msg:  message.JsonParseError,
				Data: err.Error(),
			})
			return
		}

		imsdk := dbCtx.QueryImsdk(username, params.Name)

		if imsdk.ImsdkName != params.Name {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NoSuchImsdk,
				Msg:  message.NoSuchImsdk,
				Data: nil,
			})
			return
		}

		if imsdk.ImsdkToken != token {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.TokenInvalid,
				Msg:  message.TokenInvalid,
				Data: nil,
			})
			return
		}

		cfg := dbCtx.QueryImsdkConfig(username, imsdk.ImsdkName, params.ConfigKey)

		if cfg.ConfigKey != params.ConfigKey {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NoSuchConfig,
				Msg:  message.NoSuchConfig,
				Data: nil,
			})
			return
		}

		c.JSON(constant.HttpStatusSuccess, model.Response{
			Code: code.Success,
			Msg:  message.Success,
			Data: cfg,
		})
	})

	httpServerCtx.Post(constant.PluginConfigURI, func(c *gin.Context) {
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
		var params model.ConfigParams
		err := c.ShouldBindJSON(&params)
		if err != nil {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.JsonParseError,
				Msg:  message.JsonParseError,
				Data: err.Error(),
			})
			return
		}

		plugin := dbCtx.QueryPlugin(username, params.Name)

		if plugin.PluginName != params.Name {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NoSuchPlugin,
				Msg:  message.NoSuchPlugin,
				Data: nil,
			})
			return
		}

		if plugin.PluginToken != token {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.TokenInvalid,
				Msg:  message.TokenInvalid,
				Data: nil,
			})
			return
		}

		cfg := dbCtx.QueryPluginConfig(username, plugin.PluginName, params.ConfigKey)

		if cfg.ConfigKey != params.ConfigKey {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NoSuchConfig,
				Msg:  message.NoSuchConfig,
				Data: nil,
			})
			return
		}

		c.JSON(constant.HttpStatusSuccess, model.Response{
			Code: code.Success,
			Msg:  message.Success,
			Data: cfg,
		})
	})

}

package log

import (
	"AnywhereDoorControlPlane/constant"
	"AnywhereDoorControlPlane/constant/code"
	"AnywhereDoorControlPlane/constant/message"
	"AnywhereDoorControlPlane/db"
	"AnywhereDoorControlPlane/model"
	"AnywhereDoorControlPlane/server"
	"github.com/gin-gonic/gin"
	"time"
)

type LogContext struct {
	dbCtx *db.DataBaseContext
}

func InitLogServer(dbCtx *db.DataBaseContext, httpServerCtx *server.HttpServerContext) *LogContext {
	logCtx := &LogContext{dbCtx: dbCtx}
	httpServerCtx.Post(constant.ImsdkLogURL, func(c *gin.Context) {
		logCtx.writeLogToDb(c, constant.TypeImsdk)
	})

	httpServerCtx.Post(constant.PluginLogURL, func(c *gin.Context) {
		logCtx.writeLogToDb(c, constant.TypePlugin)
	})

	return logCtx
}

func (ctx *LogContext) writeLogToDb(c *gin.Context, logType int) {
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
	var params model.LogParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(constant.HttpStatusSuccess, model.Response{
			Code: code.JsonParseError,
			Msg:  message.JsonParseError,
			Data: err.Error(),
		})
		return
	}
	var log model.Log
	log.TargetId = -1
	if logType == constant.TypeImsdk {
		imsdk := ctx.dbCtx.QueryImsdk(username, params.Name)
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
		log.TargetId = imsdk.Id
		log.Type = constant.TypeImsdk
	}
	if logType == constant.TypePlugin {
		plugin := ctx.dbCtx.QueryPlugin(username, params.Name)
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
		log.TargetId = plugin.Id
		log.Type = constant.TypePlugin
	}
	if log.TargetId == -1 {
		c.JSON(constant.HttpStatusSuccess, model.Response{
			Code: code.UnknownLogType,
			Msg:  message.UnknownLogType,
			Data: nil,
		})
		return
	}
	user := ctx.dbCtx.QueryUser(username)
	if user.Id == -1 {
		c.JSON(constant.HttpStatusSuccess, model.Response{
			Code: code.NoSuchUser,
			Msg:  message.NoSuchUser,
			Data: nil,
		})
		return
	}
	log.UserId = user.Id
	log.Timestamp = params.Timestamp
	log.Level = params.Level
	if log.Level < constant.LevelDebug {
		log.Level = constant.LevelDebug
	}
	if log.Level > constant.LevelError {
		log.Level = constant.LevelError
	}
	log.Tag = params.Tag
	log.Log = params.Log
	insertId, _ := ctx.dbCtx.InsertLog(log)
	c.JSON(constant.HttpStatusSuccess, model.Response{
		Code: code.Success,
		Msg:  message.Success,
		Data: insertId,
	})
}

func (ctx *LogContext) saveGlobalLog(tag, msg string, level int) (int, error) {
	if level < constant.LevelDebug {
		level = constant.LevelDebug
	}
	if level > constant.LevelError {
		level = constant.LevelError
	}
	return ctx.dbCtx.InsertLog(model.Log{
		Tag:       tag,
		Log:       msg,
		UserId:    constant.GlobalLogUserId,
		Type:      constant.TypeControlPlane,
		Level:     level,
		TargetId:  0,
		Timestamp: time.Now().UnixMilli(),
	})
}

func (ctx *LogContext) D(tag, msg string) {
	ctx.saveGlobalLog(tag, msg, constant.LevelDebug)
}

func (ctx *LogContext) I(tag, msg string) {
	ctx.saveGlobalLog(tag, msg, constant.LevelInfo)
}

func (ctx *LogContext) W(tag, msg string) {
	ctx.saveGlobalLog(tag, msg, constant.LevelWarn)
}

func (ctx *LogContext) E(tag, msg string) {
	ctx.saveGlobalLog(tag, msg, constant.LevelError)
}

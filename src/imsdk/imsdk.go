package imsdk

import (
	"AnywhereDoorControlPlane/apis"
	"AnywhereDoorControlPlane/constant"
	"AnywhereDoorControlPlane/constant/code"
	"AnywhereDoorControlPlane/constant/message"
	"AnywhereDoorControlPlane/db"
	"AnywhereDoorControlPlane/model"
	"AnywhereDoorControlPlane/server"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type ImsdkContext struct {
	dbCtx         *db.DataBaseContext
	httpServerCtx *server.HttpServerContext
}

func CreateImsdkContext(logCtx *apis.LogContext, dbCtx *db.DataBaseContext, httpServerCtx *server.HttpServerContext, callback func(source model.Imsdk, target model.Plugin, data string)) *ImsdkContext {
	TAG := "ImsdkContext"
	httpServerCtx.Post(constant.ImsdkURI, func(c *gin.Context) {
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
		var result model.Result
		err := c.ShouldBindJSON(&result)
		if err != nil {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.JsonParseError,
				Msg:  message.JsonParseError,
				Data: err.Error(),
			})
			logCtx.E(TAG, "bind json err. url: "+constant.ImsdkURI+", error: "+err.Error())
			return
		}

		imsdk := dbCtx.QueryImsdk(username, result.Name)

		if imsdk.ImsdkName != result.Name {
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

		plugin := dbCtx.QueryPlugin(username, result.Target)

		if plugin.PluginName != result.Target {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NoSuchPlugin,
				Msg:  message.NoSuchPlugin,
				Data: nil,
			})
			return
		}

		go callback(imsdk, plugin, result.Data)

		c.JSON(constant.HttpStatusSuccess, model.Response{
			Code: code.Success,
			Msg:  message.Success,
			Data: nil,
		})
	})

	return &ImsdkContext{
		dbCtx:         dbCtx,
		httpServerCtx: httpServerCtx,
	}
}

func (ctx *ImsdkContext) Request(logCtx *apis.LogContext, source model.Plugin, target model.Imsdk, data string) {
	TAG := "ImsdkRequest"
	url := fmt.Sprintf(constant.ImsdkURLTemplate, target.ImsdkHost, target.ImsdkPort, target.ImsdkPrefix)
	commData := model.Result{Data: data, Name: source.PluginName, Target: target.ImsdkName}
	v, err := json.Marshal(commData)
	if err != nil {
		logCtx.E(TAG, "json err. error: "+err.Error())
		return
	}
	req, err := http.NewRequest(constant.Post, url, strings.NewReader(string(v)))
	if err != nil {
		logCtx.E(TAG, "new request err. error: "+err.Error())
		return
	}
	req.Header.Add(constant.Token, target.ImsdkToken)
	req.Header.Add(constant.ContentType, constant.ContentTypeJSON)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logCtx.E(TAG, "do request err. error: "+err.Error())
	} else {
		_, err := io.ReadAll(resp.Body)
		if err != nil {
			logCtx.E(TAG, "read all resp err. error: "+err.Error())
		}
		err = resp.Body.Close()
		if err != nil {
			logCtx.E(TAG, "close resp body err. error: "+err.Error())
		}
	}
}

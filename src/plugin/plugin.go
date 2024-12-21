package plugin

import (
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

type PluginContext struct {
	dbCtx         *db.DataBaseContext
	httpServerCtx *server.HttpServerContext
}

func CreatePluginContext(dbCtx *db.DataBaseContext, httpServerCtx *server.HttpServerContext, callback func(source model.Plugin, target model.Imsdk, data string)) *PluginContext {
	httpServerCtx.Post(constant.PluginURI, func(c *gin.Context) {
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
			fmt.Println(err)
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.JsonParseError,
				Msg:  message.JsonParseError,
				Data: err.Error(),
			})
			return
		}

		plugin := dbCtx.QueryPlugin(username, result.Name)

		if plugin.PluginName != result.Name {
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

		imsdk := dbCtx.QueryImsdk(username, result.Target)

		if imsdk.ImsdkName != result.Target {
			c.JSON(constant.HttpStatusSuccess, model.Response{
				Code: code.NoSuchImsdk,
				Msg:  message.NoSuchImsdk,
				Data: nil,
			})
			return
		}

		go callback(plugin, imsdk, result.Data)

		c.JSON(constant.HttpStatusSuccess, model.Response{
			Code: code.Success,
			Msg:  message.Success,
			Data: nil,
		})
	})

	return &PluginContext{
		dbCtx:         dbCtx,
		httpServerCtx: httpServerCtx,
	}
}

func (ctx *PluginContext) Request(source model.Imsdk, target model.Plugin, data string) {
	url := fmt.Sprintf(constant.PluginURLTemplate, target.PluginHost, target.PluginPort, target.PluginPrefix)
	commData := model.Result{Data: data, Name: source.ImsdkName, Target: target.PluginName}
	v, err := json.Marshal(commData)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(constant.Post, url, strings.NewReader(string(v)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add(constant.Token, target.PluginToken)
	req.Header.Add(constant.ContentType, constant.ContentTypeJSON)
	fmt.Println(url, string(v), data)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}

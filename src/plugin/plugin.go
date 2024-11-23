package plugin

import (
	"AnywhereDoorControlPlane/db"
	"AnywhereDoorControlPlane/model"
	"AnywhereDoorControlPlane/rpc"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type PluginContext struct {
	dbCtx         *db.DataBaseContext
	httpServerCtx *rpc.HttpServerContext
}

func CreatePluginContext(dbCtx *db.DataBaseContext, httpServerCtx *rpc.HttpServerContext, callback func(source model.Plugin, target model.Imsdk, data string)) *PluginContext {
	httpServerCtx.Post("/plugin", func(c *gin.Context) {
		username := c.Request.Header.Get("username")
		token := c.Request.Header.Get("token")
		var commonData model.CommonData
		err := c.ShouldBindJSON(&commonData)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, map[string]any{
				"code":    500,
				"message": err.Error(),
			})
			return
		}

		plugin := dbCtx.QueryPlugin(username, commonData.Name)

		if plugin.PluginName != commonData.Name {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": "no such plugin",
			})
			return
		}

		if plugin.PluginToken != token {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": "token invalid",
			})
			return
		}

		imsdk := dbCtx.QueryImsdk(username, commonData.Target)

		if imsdk.ImsdkName != commonData.Target {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": "no such imsdk",
			})
			return
		}

		go callback(plugin, imsdk, commonData.Data)

		c.JSON(200, map[string]any{
			"code":    200,
			"message": "success",
		})
	})

	return &PluginContext{
		dbCtx:         dbCtx,
		httpServerCtx: httpServerCtx,
	}
}

func (ctx *PluginContext) Request(source model.Imsdk, target model.Plugin, data string) {
	url := fmt.Sprintf("%s:%d%s/plugin", target.PluginHost, target.PluginPort, target.PluginPrefix)
	commData := model.CommonData{Data: data, Name: source.ImsdkName, Target: target.PluginName}
	v, _ := json.Marshal(commData)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(v)))
	req.Header.Add("token", target.PluginToken)
	fmt.Println(url, string(v), data)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("resp = ", string(body))
	}
}

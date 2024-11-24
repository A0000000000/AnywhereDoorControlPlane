package imsdk

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

type ImsdkContext struct {
	dbCtx         *db.DataBaseContext
	httpServerCtx *rpc.HttpServerContext
}

func CreateImsdkContext(dbCtx *db.DataBaseContext, httpServerCtx *rpc.HttpServerContext, callback func(source model.Imsdk, target model.Plugin, data string)) *ImsdkContext {
	httpServerCtx.Post("/imsdk", func(c *gin.Context) {
		username := c.Request.Header.Get("username")
		token := c.Request.Header.Get("token")
		if len(username) == 0 || len(token) == 0 {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": "username or token is empty",
			})
			return
		}
		var commonData model.CommonData
		err := c.ShouldBindJSON(&commonData)
		if err != nil {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": err.Error(),
			})
			return
		}

		imsdk := dbCtx.QueryImsdk(username, commonData.Name)

		if imsdk.ImsdkName != commonData.Name {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": "no such imsdk",
			})
			return
		}

		if imsdk.ImsdkToken != token {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": "token invalid",
			})
			return
		}

		plugin := dbCtx.QueryPlugin(username, commonData.Target)

		if plugin.PluginName != commonData.Target {
			c.JSON(200, map[string]any{
				"code":    500,
				"message": "no such plugin",
			})
			return
		}

		go callback(imsdk, plugin, commonData.Data)

		c.JSON(200, map[string]any{
			"code":    200,
			"message": "success",
		})
	})

	return &ImsdkContext{
		dbCtx:         dbCtx,
		httpServerCtx: httpServerCtx,
	}
}

func (ctx *ImsdkContext) Request(source model.Plugin, target model.Imsdk, data string) {
	url := fmt.Sprintf("http://%s:%d%s/imsdk", target.ImsdkHost, target.ImsdkPort, target.ImsdkPrefix)
	commData := model.CommonData{Data: data, Name: source.PluginName, Target: target.ImsdkName}
	v, err := json.Marshal(commData)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(v)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("token", target.ImsdkToken)
	req.Header.Add("content-type", "application/json")
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

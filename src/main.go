package main

import (
	"AnywhereDoorControlPlane/apis"
	"AnywhereDoorControlPlane/db"
	"AnywhereDoorControlPlane/imsdk"
	"AnywhereDoorControlPlane/model"
	"AnywhereDoorControlPlane/plugin"
	"AnywhereDoorControlPlane/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dbCtx := db.CreateDataBaseContext()
	httpServerCtx := server.CreateHttpServer()
	var imsdkCtx *imsdk.ImsdkContext = nil
	var pluginCtx *plugin.PluginContext = nil

	logCtx := apis.InitLogServer(dbCtx, httpServerCtx)
	apis.InitConfigServer(logCtx, dbCtx, httpServerCtx)
	apis.InitRegisterServer(logCtx, dbCtx, httpServerCtx)

	imsdkCtx = imsdk.CreateImsdkContext(logCtx, dbCtx, httpServerCtx, func(source model.Imsdk, target model.Plugin, data string) {
		if pluginCtx != nil {
			pluginCtx.Request(logCtx, source, target, data)
		}
	})

	pluginCtx = plugin.CreatePluginContext(logCtx, dbCtx, httpServerCtx, func(source model.Plugin, target model.Imsdk, data string) {
		if imsdkCtx != nil {
			imsdkCtx.Request(logCtx, source, target, data)
		}
	})

	sig := make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP)
	<-sig
}

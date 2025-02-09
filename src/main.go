package main

import (
	"AnywhereDoorControlPlane/config"
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

	config.InitConfigServer(dbCtx, httpServerCtx)

	imsdkCtx = imsdk.CreateImsdkContext(dbCtx, httpServerCtx, func(source model.Imsdk, target model.Plugin, data string) {
		if pluginCtx != nil {
			pluginCtx.Request(source, target, data)
		}
	})

	pluginCtx = plugin.CreatePluginContext(dbCtx, httpServerCtx, func(source model.Plugin, target model.Imsdk, data string) {
		if imsdkCtx != nil {
			imsdkCtx.Request(source, target, data)
		}
	})

	sig := make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP)
	<-sig
}

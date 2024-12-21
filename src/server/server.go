package server

import (
	"AnywhereDoorControlPlane/constant/server"
	"github.com/gin-gonic/gin"
)

type HttpServerContext struct {
	router *gin.Engine
}

func CreateHttpServer() *HttpServerContext {
	router := gin.Default()
	go func() {
		err := router.Run(server.Port)
		if err != nil {
			panic(err)
		}
	}()
	return &HttpServerContext{router: router}
}

func (ctx *HttpServerContext) Get(url string, handler gin.HandlerFunc) {
	ctx.router.GET(url, handler)
}

func (ctx *HttpServerContext) Post(url string, handler gin.HandlerFunc) {
	ctx.router.POST(url, handler)
}

func (ctx *HttpServerContext) Put(url string, handler gin.HandlerFunc) {
	ctx.router.PUT(url, handler)
}

func (ctx *HttpServerContext) Delete(url string, handler gin.HandlerFunc) {
	ctx.router.DELETE(url, handler)
}

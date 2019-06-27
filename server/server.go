/*
@File: server.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-21 08:37  Lucien     1.0         Init
*/
package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(cors.Default())
	router.GET("/:token", get)
	router.POST("/", permanent)
	router.POST("/once", readOnce)
	router.PUT("/:key", temporary)
	router.NoRoute(notFound)
}

func Run(address string, port uint16) {
	if err := router.Run(fmt.Sprintf("%s:%d", address, port)); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"parsons.com/fds/goserver/routes"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	store := cookie.NewStore([]byte("fds_secret"))
	router.Use(sessions.Sessions("fds_remember_me", store))
	router.POST("/api", routes.HandlePost)
	router.GET("/api", routes.HandleGet)
	router.Run("localhost:8080")
}

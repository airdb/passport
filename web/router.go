package web

import (
	"github.com/airdb/passport/web/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	auth := router.Group("/")

	auth.GET("/", handlers.Auth)
	auth.HEAD("/", handlers.Auth)

	v1API := router.Group("/auth/v1")
	wechatAPI := v1API.Group("/wechat")
	wechatAPI.POST("/login", handlers.WechatLogin)
	wechatAPI.GET("/login", handlers.WechatLogin)
	wechatAPI.GET("/", handlers.Auth)
	wechatAPI.HEAD("/", handlers.Auth)

	wechatAPI.GET("/logout", handlers.WechatLogout)
	return &Router{
		router,
	}
}

func (r *Router) Run() {
	err := r.Engine.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("start failed")
	}
}

func Run() {
	NewRouter().Run()
}

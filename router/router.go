package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe_learn/app/api"
	"goframe_learn/app/middleware"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.ALL("/", api.Hello)
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.POST("/login", middleware.Auth.LoginHandler)
				group.POST("/", api.User.SignUp)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin)
					group.GET("/", api.User.GetInfo)
				})
			})
		})
	})
}

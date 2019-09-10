package router

import (
	"github.com/gin-gonic/gin"
	"lilith/controller"
	"lilith/middleware"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/go", controller.NavGo)
		v1.Use(middleware.OnlineMiddleware())
		//导航相关接口
		v1.GET("/navs/type/:type", controller.NavList)
		v1.GET("/hots", controller.NavHot)
		v1.POST("/navs", controller.CreateNav)
		v1.GET("/navs/title/:title", controller.FindByTitle)
		v1.POST("/likes/:id", controller.NavLike)
		//登录注册
		v1.POST("/login", controller.Login)
		v1.POST("/register", controller.Register)

		v1.GET("/fishes/:type", controller.FishAll)

		v1.GET("/resources", controller.ResourceList)
		v1.GET("/onlines", controller.OnlineUsersNo)
		v1.GET("/nodes", controller.ListByNode)
		v1.GET("/bookmarks", controller.ListBookmark)

		//以下接口须认证
		v1.Use(middleware.AuthJwt())
		v1.POST("/resources", controller.CreateResource)
		v1.GET("/resources/:id", controller.ResourceInfo)
		v1.POST("/comments", controller.CreateComment)
		v1.GET("/comments/:id/:page", controller.CommentList)
		v1.POST("/communications", controller.CreateCommunication)
		v1.GET("/communications/:id", controller.CommunicationInfo)
		v1.POST("/feedbacks", controller.CreateFeedback)
		v1.POST("/bookmarks", controller.CreateBookmark)
		v1.GET("/collections", controller.CollectionAll)
		v1.POST("/collections/:id", controller.AddCollection)
		v1.POST("/pwd", controller.ChangePwd)
		v1.POST("/users", controller.UpdateUserInfo)
	}
	return r
}

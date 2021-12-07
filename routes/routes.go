package routes

import (
	v1 "GO-GIN-Vue-blog/api/v1"
	"GO-GIN-Vue-blog/middleware"
	"GO-GIN-Vue-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppNode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		Auth.POST("category/add", v1.AddCategory)
		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DeleteCategory)

		//文章模块的路由接口
		Auth.POST("article/add", v1.AddArticle)
		Auth.PUT("article/:id", v1.EditArticle)
		Auth.DELETE("article/:id", v1.DeleteArticle)

		//文件上传
		Auth.POST("upload", v1.UpLoad)
	}
	routeV1 := r.Group("api/v1")
	{
		routeV1.POST("user/add", v1.AddUser)
		routeV1.GET("users", v1.GetUsers)
		routeV1.GET("categories", v1.GetCategories)
		//分别为查询所有文章,分类包含文章,单个文章
		routeV1.GET("articles", v1.GetArticles)
		routeV1.GET("category/articles/:cid", v1.GetCategoryAllArticles)
		routeV1.GET("category/article/:id", v1.GetArticle)
		routeV1.POST("Login", v1.Login)
	}
	r.Run(utils.HttpPort)
}

package routes

import (
	v1 "GO-GIN-Vue-blog/api/v1"
	"GO-GIN-Vue-blog/middleware"
	"GO-GIN-Vue-blog/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	//后端管理接口
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		Auth.GET("admin/users", v1.GetUsers)
		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DeleteUser)
		Auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)

		//分类模块的路由接口
		Auth.GET("admin/categories", v1.GetCategories)
		Auth.POST("category/add", v1.AddCategory)
		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DeleteCategory)

		//文章模块的路由接口
		Auth.GET("admin/category/article/:id", v1.GetArticle)
		Auth.GET("admin/articles", v1.GetArticles)
		Auth.POST("article/add", v1.AddArticle)
		Auth.PUT("article/:id", v1.EditArticle)
		Auth.DELETE("article/:id", v1.DeleteArticle)

		//文件上传模块
		Auth.POST("upload", v1.UpLoad)

		//个人信息模块
		Auth.GET("admin/profile/:id", v1.GetProfile)
		Auth.PUT("profile/:id", v1.UpdateProfile)

		//评论模块
		Auth.GET("comment/list", v1.GetCommentList)
		Auth.DELETE("delcomment/:id", v1.DeleteComment)
		Auth.PUT("checkcomment/:id", v1.CheckComment)
		Auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	//前端展示接口
	routeV1 := r.Group("api/v1")
	{
		//用户信息模块
		routeV1.POST("user/add", v1.AddUser)
		routeV1.GET("users", v1.GetUsers)
		routeV1.GET("user/:id", v1.GetUserInfo)

		//分类模块
		routeV1.GET("categories", v1.GetCategories)
		routeV1.GET("category/:id", v1.GetCateInfo)

		//文章模块
		//分别为查询所有文章,分类包含文章,单个文章
		routeV1.GET("articles", v1.GetArticles)
		routeV1.GET("category/articles/:cid", v1.GetCategoryAllArticles)
		routeV1.GET("category/article/:id", v1.GetArticle)

		//登录控制模块
		routeV1.POST("Login", v1.Login)
		routeV1.POST("loginfront", v1.LoginFront)

		//个人信息模块
		routeV1.GET("profile/:id", v1.GetProfile)

		//评论模块
		routeV1.POST("addcomment", v1.AddComment)
		routeV1.GET("comment/info/:id", v1.GetComment)
		routeV1.GET("commentfront/:id", v1.GetCommentListFront)
		routeV1.GET("commentcount/:id", v1.GetCommentCount)
	}
	r.Run(utils.HttpPort)
}

package main

import (
	"blog/handlers"
	"blog/middleware"
	"blog/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移模型
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// 创建Gin路由
	r := gin.Default()

	// 添加数据库实例到中间件 - 这是关键部分！
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // 将数据库实例设置到上下文中
		c.Next()
	})

	// 公共路由
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/posts", handlers.GetPosts)
	r.GET("/posts/:id", handlers.GetPost)
	r.GET("/posts/:id/comments", handlers.GetComments)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Blog API",
			"endpoints": map[string]string{
				"register":       "POST /register",
				"login":          "POST /login",
				"get_posts":      "GET /posts",
				"get_post":       "GET /posts/:id",
				"get_comments":   "GET /posts/:id/comments",
				"create_post":    "POST /posts (auth required)",
				"update_post":    "PUT /posts/:id (auth required)",
				"delete_post":    "DELETE /posts/:id (auth required)",
				"create_comment": "POST /posts/:id/comments (auth required)",
			},
		})
	})

	// 需要认证的路由
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/posts", handlers.CreatePost)
		auth.PUT("/posts/:id", handlers.UpdatePost)
		auth.DELETE("/posts/:id", handlers.DeletePost)
		auth.POST("/posts/:id/comments", handlers.CreateComment)
	}

	// 启动服务器
	r.Run(":8080")
}

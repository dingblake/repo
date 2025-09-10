package main

/*题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。*/

/*基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。*/

/*题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。*/
import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID            uint      `gorm:"primaryKey"`
	Username      string    `gorm:"not null;unique"`
	Email         string    `gorm:"not null;unique"`
	Password_hash string    `gorm:"not null"`
	Created_at    time.Time `gorm:"autoCreateTime"`
	PostCount     int       // 文章数量统计字段
	Posts         []Post    // 一对多关系
}

type Post struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Content       string
	UserID        uint
	CommentCount  int       `gorm:"default:0"`             // 评论数量统计字段
	CommentStatus string    `gorm:"default:'无评论';size:20"` // 评论状态
	Comments      []Comment // 一对多关系
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	Content   string
	PostID    uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func main() {
	db, err := gorm.Open(mysql.Open("admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	// 初始化账户数据
	initblogs(db)

	var posts []Post
	var post Post
	userid := 1

	// 查询某个用户发布的所有文章及其评论
	fmt.Println("\n查询用户1的所有文章及其评论:")
	db.Preload("User").Preload("Comments").Where("user_id = ?", userid).Find(&posts)
	for _, post := range posts {
		fmt.Printf("文章: %s (评论数: %d, 状态: %s)\n", post.Title, post.CommentCount, post.CommentStatus)
		for _, comment := range post.Comments {
			fmt.Printf("  - 评论: %s\n", comment.Content)
		}
	}

	// 查询评论数量最多的文章
	fmt.Println("\n查询评论数量最多的文章:")
	db.Order("comment_count DESC").First(&post)
	fmt.Printf("最受欢迎的文章: %s (评论数: %d)\n", post.Title, post.CommentCount)
	// 测试删除评论
	fmt.Println("\n测试删除评论:")
	var comment Comment
	db.First(&comment)
	db.Delete(&comment)

	// 重新查询文章状态

	db.First(&post, 1)
	fmt.Printf("删除一条评论后，文章状态: %s (评论数: %d)\n", post.CommentStatus, post.CommentCount)

	// 删除所有评论
	db.Where("post_id = ?", 1).Delete(&Comment{})

	// 重新查询文章状态
	db.First(&post, 1)
	fmt.Printf("删除所有评论后，文章状态: %s (评论数: %d)\n", post.CommentStatus, post.CommentCount)

	// 查询用户信息以验证文章计数
	var user User
	db.First(&user, 1)
	fmt.Printf("用户 %s 的文章数量: %d\n", user.Username, user.PostCount)

}

func initblogs(db *gorm.DB) error {
	var count int64
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return fmt.Errorf("检查用户数据失败: %v", err)
	}
	if err := db.Model(&Post{}).Count(&count).Error; err != nil {
		return fmt.Errorf("检查文章数据失败: %v", err)
	}
	if err := db.Model(&Comment{}).Count(&count).Error; err != nil {
		return fmt.Errorf("检查评论数据失败: %v", err)
	}
	if count > 0 {
		fmt.Println("数据库中已有数据，跳过初始化")
		return nil
	}

	users := []User{
		{Username: "张三", Email: "alice@example.com", Password_hash: "123456"},
		{Username: "王五", Email: "bob@example.com", Password_hash: "123456"},
	}
	if result := db.Create(&users); result.Error != nil {
		return fmt.Errorf("用户初始化失败:", result.Error)
	}

	posts := []Post{
		{Title: "我的第一篇文章", Content: "这是关于Go语言入门的文章内容...", UserID: users[0].ID},
		{Title: "我的第二篇文章", Content: "这是关于Gorm使用的文章内容...", UserID: users[0].ID},
		{Title: "Web开发实践", Content: "这是关于Web开发实践的文章内容...", UserID: users[1].ID},
	}

	if result := db.Create(&posts); result.Error != nil {
		return fmt.Errorf("文章初始化失败:", result.Error)
	}

	comments := []Comment{
		{Content: "很好的入门文章!", PostID: posts[0].ID},
		{Content: "期待更多关于Go的内容。", PostID: posts[0].ID},
		{Content: "Gorm确实很方便。", PostID: posts[1].ID},
	}
	if result := db.Create(&comments); result.Error != nil {
		return fmt.Errorf("评论初始化失败:", result.Error)
	}

	fmt.Println("初始化成功")
	return nil
}
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1"))
	return
}

// Comment 的 AfterCreate 钩子
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论数量
	tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count + 1"))

	// 更新文章的评论状态
	tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_status", "有评论")
	return
}

// Comment 的 AfterDelete 钩子
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 更新文章的评论数量
	tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count - 1"))

	// 检查当前文章的评论数量
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论")
	}
	return
}

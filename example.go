package main

// import (
// 	"errors"
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Account struct {
// 	AccountID string `gorm:"primaryKey;column:account_id"`
// 	Balance   int    `gorm:"column:balance"`
// }

// func (Account) TableName() string {
// 	return "accounts"
// }

// type Transaction struct {
// 	gorm.Model
// 	FromAccountID string `gorm:"column:from_account_id"`
// 	ToAccountID   string `gorm:"column:to_account_id"`
// 	Amount        int    `gorm:"column:amount"`
// }

// func (Transaction) TableName() string {
// 	return "transactions"
// }

// func main() {
// 	db, err := gorm.Open(mysql.Open("admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("数据库连接失败:", err)
// 	}

// 	// 自动迁移表结构
// 	if err := db.AutoMigrate(&Account{}, &Transaction{}); err != nil {
// 		log.Fatal("表结构迁移失败:", err)
// 	}

// 	// 初始化账户数据
// 	initAccounts(db)

// 	// 执行转账
// 	if err := transferMoney(db, "A", "B", 100); err != nil {
// 		log.Fatal("转账失败:", err)
// 	}

// 	fmt.Println("转账成功!")
// }

// func initAccounts(db *gorm.DB) {
// 	// 检查是否已初始化
// 	var count int64
// 	db.Model(&Account{}).Count(&count)
// 	if count > 0 {
// 		return
// 	}

// 	accounts := []Account{
// 		{AccountID: "A", Balance: 1200},
// 		{AccountID: "B", Balance: 0},
// 		{AccountID: "C", Balance: 0},
// 	}

// 	if result := db.Create(&accounts); result.Error != nil {
// 		log.Fatal("初始化账户失败:", result.Error)
// 	}

// 	fmt.Println("账户初始化成功")
// }

// func transferMoney(db *gorm.DB, fromID, toID string, amount int) error {
// 	// 使用事务确保操作的原子性
// 	return db.Transaction(func(tx *gorm.DB) error {
// 		// 1. 检查转出账户是否存在且余额充足
// 		var fromAccount Account
// 		if err := tx.Where("account_id = ?", fromID).First(&fromAccount).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				return fmt.Errorf("转出账户不存在")
// 			}
// 			return err
// 		}

// 		if fromAccount.Balance < amount {
// 			return fmt.Errorf("账户余额不足，当前余额: %d", fromAccount.Balance)
// 		}

// 		// 2. 检查转入账户是否存在
// 		var toAccount Account
// 		if err := tx.Where("account_id = ?", toID).First(&toAccount).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				return fmt.Errorf("转入账户不存在")
// 			}
// 			return err
// 		}

// 		// 3. 执行转账操作
// 		// 扣除转出账户余额
// 		if err := tx.Model(&Account{}).
// 			Where("account_id = ? AND balance >= ?", fromID, amount). // 添加条件防止并发问题
// 			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
// 			return fmt.Errorf("扣除转出账户余额失败: %v", err)
// 		}

// 		// 增加转入账户余额
// 		if err := tx.Model(&Account{}).
// 			Where("account_id = ?", toID).
// 			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
// 			return fmt.Errorf("增加转入账户余额失败: %v", err)
// 		}

// 		// 4. 记录交易
// 		transaction := Transaction{
// 			FromAccountID: fromID,
// 			ToAccountID:   toID,
// 			Amount:        amount,
// 		}

// 		if err := tx.Create(&transaction).Error; err != nil {
// 			return fmt.Errorf("记录交易失败: %v", err)
// 		}

// 		// 5. 验证结果
// 		var updatedFromAccount Account
// 		if err := tx.Where("account_id = ?", fromID).First(&updatedFromAccount).Error; err != nil {
// 			return err
// 		}

// 		var updatedToAccount Account
// 		if err := tx.Where("account_id = ?", toID).First(&updatedToAccount).Error; err != nil {
// 			return err
// 		}

// 		fmt.Printf("转账成功:\n")
// 		fmt.Printf("  转出账户 %s 余额: %d -> %d\n", fromID, fromAccount.Balance, updatedFromAccount.Balance)
// 		fmt.Printf("  转入账户 %s 余额: %d -> %d\n", toID, toAccount.Balance, updatedToAccount.Balance)

// 		return nil
// 	})
// }

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 模型
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;size:100"`
	Email     string `gorm:"unique;not null;size:150"`
	PostCount int    `gorm:"default:0"`         // 文章数量统计字段
	Posts     []Post `gorm:"foreignKey:UserID"` // 一对多关系
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"not null;size:200"`
	Content       string    `gorm:"type:text;not null"`
	UserID        uint      `gorm:"not null;index"`        // 外键
	CommentCount  int       `gorm:"default:0"`             // 评论数量统计字段
	CommentStatus string    `gorm:"default:'无评论';size:20"` // 评论状态
	User          User      `gorm:"foreignKey:UserID"`     // 关联用户
	Comments      []Comment `gorm:"foreignKey:PostID"`     // 一对多关系
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Comment 模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null;index"`    // 外键
	Post      Post   `gorm:"foreignKey:PostID"` // 关联文章
	CreatedAt time.Time
	UpdatedAt time.Time
}

var db *gorm.DB

func initDB() {
	// MySQL连接配置
	dsn := "admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		), // 自定义日志输出
	})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// 自动迁移表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	fmt.Println("Database initialized successfully")
}

// 创建示例数据
func createSampleData() {
	// 创建用户
	user := User{Name: "张三", Email: "zhangsan@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Error creating user: ", result.Error)
	}

	// 创建文章
	post := Post{
		Title:   "我的第一篇文章",
		Content: "这是文章内容...",
		UserID:  user.ID,
	}
	result = db.Create(&post)
	if result.Error != nil {
		log.Fatal("Error creating post: ", result.Error)
	}

	// 创建评论
	comments := []Comment{
		{Content: "好文章!", PostID: post.ID},
		{Content: "写得很棒!", PostID: post.ID},
	}
	result = db.Create(&comments)
	if result.Error != nil {
		log.Fatal("Error creating comments: ", result.Error)
	}

	fmt.Println("Sample data created successfully")
}

// 查询某个用户发布的所有文章及其对应的评论信息
func getUserPostsWithComments(userID uint) ([]Post, error) {
	var posts []Post
	result := db.Preload("User").Preload("Comments").Where("user_id = ?", userID).Find(&posts)
	return posts, result.Error
}

// 查询评论数量最多的文章信息
func getMostCommentedPost() (Post, error) {
	var post Post
	result := db.Order("comment_count DESC").First(&post)
	return post, result.Error
}

// Post 的 AfterCreate 钩子
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

func main() {
	// 初始化数据库
	initDB()

	// 创建示例数据
	createSampleData()

	// 查询用户1的所有文章及其评论
	fmt.Println("\n查询用户1的所有文章及其评论:")
	posts, err := getUserPostsWithComments(1)
	if err != nil {
		log.Fatal("Error querying user posts: ", err)
	}

	for _, post := range posts {
		fmt.Printf("文章: %s (评论数: %d, 状态: %s)\n", post.Title, post.CommentCount, post.CommentStatus)
		for _, comment := range post.Comments {
			fmt.Printf("  - 评论: %s\n", comment.Content)
		}
	}

	// 查询评论数量最多的文章
	fmt.Println("\n查询评论数量最多的文章:")
	mostCommentedPost, err := getMostCommentedPost()
	if err != nil {
		log.Fatal("Error querying most commented post: ", err)
	}
	fmt.Printf("最受欢迎的文章: %s (评论数: %d)\n", mostCommentedPost.Title, mostCommentedPost.CommentCount)

	// 测试删除评论
	fmt.Println("\n测试删除评论:")
	var comment Comment
	db.First(&comment)
	db.Delete(&comment)

	// 重新查询文章状态
	var post Post
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
	fmt.Printf("用户 %s 的文章数量: %d\n", user.Name, user.PostCount)
}

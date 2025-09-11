package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Posts    []Post    `json:"-"`
	Comments []Comment `json:"-"`
}

type Post struct {
	gorm.Model
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null;type:text" json:"content"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"-"`
	Comments  []Comment `json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	gorm.Model
	Content   string    `gorm:"not null" json:"content"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	PostID    uint      `json:"post_id"`
	Post      Post      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	dsn := "admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

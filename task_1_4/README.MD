# 个人博客系统后端

基于 Go + Gin + GORM 开发的个人博客系统后端，支持用户认证、文章管理和评论功能。

## 功能特性

- 用户注册和登录（JWT认证）
- 文章的创建、读取、更新和删除（CRUD）
- 评论的创建和读取
- 权限控制（用户只能操作自己的文章）

## 技术栈

- Go 1.21+
- Gin Web框架
- GORM ORM库
- MySQL数据库
- JWT认证

## 安装和运行

1. 确保已安装Go 1.21+和MySQL

2. 克隆项目并安装依赖：
```bash
cd /usr/gopro/gotest/init_project/task_1_4
go mod tidy
3. 创建MySQL数据库：
4. 修改数据库连接配置（在models/models.go中）：
    dsn := "用户名:密码@tcp(数据库地址:端口)/blog?charset=utf8mb4&parseTime=True&loc=Local"
5. 运行应用：go run main.go
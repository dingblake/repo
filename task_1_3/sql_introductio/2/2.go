package main

/*题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。*/
import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Books struct {
	ID     string  `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	// 连接数据库
	db, err := sqlx.Connect("mysql", "admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()
	//初始化数据
	if err := initBooks(db); err != nil {
		log.Fatal("初始化失败:", err)
	}

	// 1. 查询价格大于50的书籍
	var book []Books
	err = db.Select(&book, "SELECT * FROM books WHERE Price > ?", 50)
	if err != nil {
		log.Fatal("查询失败:", err)
	}

	fmt.Println("技术部的员工信息:")
	for _, i := range book {
		fmt.Printf("书籍ID: %s, 书籍: %s, 作者: %s ,价格: %f\n", i.ID, i.Title, i.Author, i.Price)
	}

}

func initBooks(db *sqlx.DB) error {
	// 1. 检查表是否存在，如果不存在则创建
	var tableExists bool
	err := db.Get(&tableExists, `
        SELECT COUNT(*) > 0 
        FROM information_schema.tables 
        WHERE table_schema = DATABASE() AND table_name = 'books'
    `)
	if err != nil {
		return fmt.Errorf("检查表是否存在失败: %v", err)
	}

	if !tableExists {
		// 创建 books 表
		createTableSQL := `
            CREATE TABLE books (
                id VARCHAR(50) PRIMARY KEY,
                title VARCHAR(1000) NOT NULL,
                author VARCHAR(100),
                price DECIMAL(10,2)
            )
        `
		_, err := db.Exec(createTableSQL)
		if err != nil {
			return fmt.Errorf("创建表失败: %v", err)
		}
		fmt.Println("books 表创建成功")
	} else {
		fmt.Println("books 表已存在")
	}

	// 2. 检查表中是否有数据
	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM books")
	if err != nil {
		return fmt.Errorf("查询数据数量失败: %v", err)
	}

	if count > 0 {
		fmt.Println("表中已有数据，跳过初始化")
		return nil
	}

	// 3. 使用事务插入初始化数据
	ctx := context.Background()
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("开始事务失败: %v", err)
	}

	// 确保在函数返回时处理事务
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // 重新抛出panic
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				fmt.Printf("提交事务失败: %v\n", err)
			}
		}
	}()

	// 准备插入数据的SQL语句
	insertSQL := `INSERT INTO books (id, title, author, price) VALUES (?, ?, ?, ?)`

	// 定义要插入的数据
	books := []struct {
		ID     string
		Title  string
		Author string
		Price  float64
	}{
		{"001", "三体", "刘慈欣", 45.90},
		{"002", "活着", "余华", 28.00},
		{"003", "红楼梦", "曹雪芹", 59.80},
		{"004", "平凡的世界", "路遥", 39.90},
	}

	// 执行插入操作
	for _, book := range books {
		_, err = tx.ExecContext(ctx, insertSQL, book.ID, book.Title, book.Author, book.Price)
		if err != nil {
			return fmt.Errorf("插入数据失败: %v", err)
		}
	}

	fmt.Printf("成功插入 %d 条数据\n", len(books))
	return nil
}

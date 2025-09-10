package main

/*题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。*/
import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func main() {
	// 连接数据库
	db, err := sqlx.Connect("mysql", "admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()
	//初始化数据
	if err := initEmployees(db); err != nil {
		log.Fatal("初始化失败:", err)
	}
	// 1. 查询所有部门为 "技术部" 的员工
	var techEmployees []Employee
	err = db.Select(&techEmployees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal("查询失败:", err)
	}

	fmt.Println("技术部的员工信息:")
	for _, emp := range techEmployees {
		fmt.Printf("员工ID: %s, 姓名: %s, 薪资: %d\n", emp.ID, emp.Name, emp.Salary)
	}

	// 2. 查询工资最高的员工
	var maxSalaryEmployee Employee
	err = db.Get(&maxSalaryEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal("查询失败:", err)
	}

	fmt.Printf("\n工资最高的员工信息: 员工ID: %s, 姓名: %s, 部门: %s, 薪资: %d\n",
		maxSalaryEmployee.ID, maxSalaryEmployee.Name, maxSalaryEmployee.Department, maxSalaryEmployee.Salary)
}
func initEmployees(db *sqlx.DB) error {
	// 1. 检查表是否存在，如果不存在则创建
	var tableExists bool
	err := db.Get(&tableExists, `
        SELECT COUNT(*) > 0 
        FROM information_schema.tables 
        WHERE table_schema = DATABASE() AND table_name = 'employees'
    `)
	if err != nil {
		return fmt.Errorf("检查表是否存在失败: %v", err)
	}

	if !tableExists {
		// 创建 employees 表
		createTableSQL := `
            CREATE TABLE employees (
                id VARCHAR(50) PRIMARY KEY,
                name VARCHAR(100) NOT NULL,
                department VARCHAR(100),
                salary INT
            )
        `
		_, err := db.Exec(createTableSQL)
		if err != nil {
			return fmt.Errorf("创建表失败: %v", err)
		}
		fmt.Println("employees 表创建成功")
	} else {
		fmt.Println("employees 表已存在")
	}

	// 2. 检查表中是否有数据
	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM employees")
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
	insertSQL := `INSERT INTO employees (id, name, department, salary) VALUES (?, ?, ?, ?)`

	// 定义要插入的数据
	employees := []struct {
		ID         string
		Name       string
		Department string
		Salary     int
	}{
		{"001", "张三", "技术部", 1000},
		{"002", "李四", "技术部", 1200},
		{"003", "王五", "财务部", 1000},
		{"004", "赵四", "主管部", 12000},
	}

	// 执行插入操作
	for _, emp := range employees {
		_, err = tx.ExecContext(ctx, insertSQL, emp.ID, emp.Name, emp.Department, emp.Salary)
		if err != nil {
			return fmt.Errorf("插入数据失败: %v", err)
		}
	}

	fmt.Printf("成功插入 %d 条员工记录\n", len(employees))
	return nil
}

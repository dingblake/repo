package main

/*题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。*/

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open("admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	type Student struct {
		gorm.Model
		Name  string
		Age   int
		Grade string
	}
	db.AutoMigrate(&Student{})
	student1 := []Student{{Name: "张三", Age: 20, Grade: "三年级"},
		{Name: "李四", Age: 18, Grade: "二年级"},
		{Name: "王五", Age: 19, Grade: "三年级"},
		{Name: "丁已", Age: 14, Grade: "二年级"},
		{Name: "赵四", Age: 19, Grade: "一年级"}}
	result := db.Create(&student1) // pass a slice to insert multiple row
	if result.Error != nil {
		panic(result.Error)
	}
	// 输出插入的行数
	fmt.Printf("成功插入 %d 条记录\n", result.RowsAffected)

	//var upstruct Student
	var foundStudent Student
	result = db.Where("age > ?", 18).First(&foundStudent)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 处理记录不存在的情况
			fmt.Println("没有找到符合条件的记录")
		} else {
			// 处理其他错误
			panic(result.Error)
		}
	} else {
		// 处理找到记录的情况
		fmt.Printf("找到学生: %s, 年龄: %d\n", foundStudent.Name, foundStudent.Age)
	}

	result = db.Model(&Student{}).Where("name = ?", "张三").Update("Grade", "四年级")
	fmt.Printf("更新了 %d 条记录\n", result.RowsAffected)
	db.Where("age < ?", 15).Delete(&Student{})
	result = db.Unscoped().Where("age < ?", 15).Delete(&Student{})

	fmt.Printf("删除了 %d 条记录\n", result.RowsAffected)

	// 验证删除结果
	var remainingStudents []Student
	db.Find(&remainingStudents)
	fmt.Println("剩余学生:")
	for _, student := range remainingStudents {
		fmt.Printf("姓名: %s, 年龄: %d,年级：%s\n", student.Name, student.Age, student.Grade)
	}

}

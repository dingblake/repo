package main

/*假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。*/
import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Accounts struct {
	Account_id string `gorm:"primaryKey"`
	Balance    int
}
type Transactions struct {
	gorm.Model
	From_account_id string
	To_account_id   string
	Amount          int
}

func main() {

	db, err := gorm.Open(mysql.Open("admin:123456@tcp(127.0.0.1:3306)/task_data?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var result *gorm.DB

	// 初始化账户数据
	initAccounts(db)

	balances := 100
	var idsa Accounts
	var idsb Accounts
	var idsau Accounts
	var idsbu Accounts

	result = db.Where("account_id = ? AND balance > ? ", "A", balances).First(&idsa)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 处理记录不存在的情况
			fmt.Println("转账失败，没找到账户或该账户下余额不足100")

		} else {
			// 处理其他错误
			panic(result.Error)
		}
	} else {
		// 处理找到记录的情况
		fmt.Printf("找到账户: %s, 余额: %d\n", idsa.Account_id, idsa.Balance)
		db.Where("account_id = ? ", "B").First(&idsb)

		db.Model(&Accounts{}).Where("account_id = ?", "A").Update("Balance", idsa.Balance-balances)
		db.Model(&Accounts{}).Where("account_id = ?", "B").Update("Balance", idsb.Balance+balances)
		//fmt.Printf("更新了 %d 条记录\n", result.RowsAffected)
		db.Where("account_id = ? ", "A").First(&idsau)
		fmt.Printf("账户: %s, 余额: %d\n", idsau.Account_id, idsau.Balance)
		db.Where("account_id = ? ", "B").First(&idsbu)
		fmt.Printf("账户: %s, 余额: %d\n", idsbu.Account_id, idsbu.Balance)
		trans := Transactions{From_account_id: idsa.Account_id, To_account_id: idsb.Account_id, Amount: balances}
		result = db.Create(&trans)
		if result.Error != nil {
			panic(result.Error)
		}
		//fmt.Printf("成功插入 %d 条记录\n", result.RowsAffected)
		fmt.Println("转账成功!")
		fmt.Printf("转出账户 %s 余额: %d -> %d\n", idsau.Account_id, idsa.Balance, idsau.Balance)
		fmt.Printf("转入账户 %s 余额: %d -> %d\n", idsbu.Account_id, idsb.Balance, idsbu.Balance)
	}

}

func initAccounts(db *gorm.DB) {
	// 检查是否已初始化
	var count int64
	db.Model(&Accounts{}).Count(&count)
	if count > 0 {
		return
	}

	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transactions{})
	accounts := []Accounts{
		{Account_id: "A", Balance: 1200},
		{Account_id: "B", Balance: 0},
		{Account_id: "C", Balance: 0},
	}

	if result := db.Create(&accounts); result.Error != nil {
		log.Fatal("初始化账户失败:", result.Error)
	}

	fmt.Println("账户初始化成功")
}

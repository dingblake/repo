package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/test/init_project/task_4_2/counter" // 替换为你的模块名
)

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("警告: 无法加载 .env 文件")
	}
}

// 验证 Infura URL
func validateInfuraURL(apiKey string) (string, error) {
	if apiKey == "" {
		return "", fmt.Errorf("INFURA_API_KEY 为空")
	}

	// 清理可能的空格或引号
	apiKey = strings.TrimSpace(apiKey)
	apiKey = strings.Trim(apiKey, `"'`)

	// 检查长度（Infura Project ID 通常是32位）
	if len(apiKey) != 32 {
		return "", fmt.Errorf("INFURA_API_KEY 长度异常，应为32位，实际为%d位", len(apiKey))
	}

	return fmt.Sprintf("https://sepolia.infura.io/v3/%s", apiKey), nil
}
func main() {
	// 连接到 Sepolia 测试网络
	loadEnv()

	// 获取并验证 API Key
	apiKey := os.Getenv("INFURA_API_KEY")
	if apiKey == "" {
		log.Fatal("错误: 请设置 INFURA_API_KEY 环境变量")
	}

	infuraURL, err := validateInfuraURL(apiKey)
	if err != nil {
		log.Fatalf("INFURA_API_KEY 验证失败: %v", err)
	}
	// 配置

	//infuraURL = fmt.Sprintf("https://sepolia.infura.io/v3/%s", os.Getenv("INFURA_API_KEY"))
	privateKeyStr := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 从私钥字符串解析出私钥对象
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	// 从私钥推导出公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 从公钥推导出地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 创建认证器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// 部署合约或连接到已部署的合约
	var contract *counter.Counter
	contractAddress := common.HexToAddress("0x0Fa97A70A67A77BeC14A9269dEd138ad3Ef61c63") // 已部署的合约地址

	if contractAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		// 部署新合约
		fmt.Println("Deploying new contract...")
		auth.Value = big.NewInt(0) // 设置交易参数
		auth.GasLimit = 3000000

		address, tx, instance, err := counter.DeployCounter(auth, client)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Contract deployed! Address: %s\n", address.Hex())
		fmt.Printf("Transaction: %s\n", tx.Hash().Hex())

		contract = instance
		contractAddress = address

		// 等待部署确认
		time.Sleep(30 * time.Second)
	} else {
		// 连接到已部署的合约
		fmt.Printf("Connecting to existing contract: %s\n", contractAddress.Hex())
		contract, err = counter.NewCounter(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 交互示例
	interactWithContract(client, contract, auth, fromAddress)
}

func interactWithContract(client *ethclient.Client, contract *counter.Counter, auth *bind.TransactOpts, fromAddress common.Address) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel() // 确保释放资源
	// 1. 读取当前计数
	currentCount, err := contract.GetCount(&bind.CallOpts{
		From:    fromAddress,
		Context: ctx,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current count: %d\n", currentCount)

	// 2. 增加计数
	fmt.Println("Incrementing count...")
	auth.Value = big.NewInt(0)
	auth.GasLimit = 300000

	tx, err := contract.Increment(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// 等待交易确认
	fmt.Println("Waiting for transaction confirmation...")
	time.Sleep(30 * time.Second)

	// 3. 再次读取计数
	newCount, err := contract.GetCount(&bind.CallOpts{
		From: fromAddress,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New count: %d\n", newCount)

	// 4. 监听事件
	fmt.Println("Setting up event listener...")
	watchEvents(contract, fromAddress)
}

func watchEvents(contract *counter.Counter, fromAddress common.Address) {
	events := make(chan *counter.CounterCountIncreased)
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	subscription, err := contract.WatchCountIncreased(opts, events)
	if err != nil {
		log.Fatal(err)
	}
	defer subscription.Unsubscribe()

	fmt.Println("Listening for CountIncreased events...")

	for {
		select {
		case err := <-subscription.Err():
			log.Fatal(err)
		case event := <-events:
			fmt.Printf("Event received - New Count: %d\n", event.NewCount)
		case <-time.After(60 * time.Second):
			fmt.Println("Event listening timeout")
			return
		}
	}
}

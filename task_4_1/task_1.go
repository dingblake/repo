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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type BlockchainClient struct {
	client *ethclient.Client
}

func NewBlockchainClient(infuraURL string) (*BlockchainClient, error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, err
	}
	return &BlockchainClient{client: client}, nil
}

func (bc *BlockchainClient) Close() {
	bc.client.Close()
}

// 查询区块信息
func (bc *BlockchainClient) QueryBlock(blockNumber int64) error {
	blockNum := big.NewInt(blockNumber)
	block, err := bc.client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		return err
	}

	fmt.Println("=== 区块详细信息 ===")
	fmt.Printf("区块号: %d\n", block.Number().Int64())
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("区块时间戳: %d (%s)\n", block.Time(), time.Unix(int64(block.Time()), 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	fmt.Printf("矿工地址: %s\n", block.Coinbase().Hex())
	fmt.Printf("Gas 使用量/限制: %d / %d\n", block.GasUsed(), block.GasLimit())
	fmt.Print("==================\n")

	return nil
}

// 发送交易
func (bc *BlockchainClient) SendTransaction(privateKeyHex, toAddress string, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	nonce, err := bc.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := bc.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	chainID, err := bc.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	to := common.HexToAddress(toAddress)

	tx := types.NewTransaction(
		nonce,
		to,
		amount,
		21000,
		gasPrice,
		nil,
	)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = bc.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// 查询余额
func (bc *BlockchainClient) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	return bc.client.BalanceAt(context.Background(), account, nil)
}

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

	privateKey := os.Getenv("PRIVATE_KEY")

	// 创建客户端
	client, err := NewBlockchainClient(infuraURL)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	fmt.Print("=== Sepolia 测试网络交互演示 ===\n")
	//fmt.Print(infuraURL)

	// 1. 查询最新区块信息
	fmt.Println("1. 查询最新区块信息:")
	header, err := client.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("获取最新区块头失败: %v", err)
	}

	latestBlock := header.Number.Int64()
	fmt.Printf("当前最新区块: %d\n", latestBlock)

	// 查询最新区块和前一个区块
	client.QueryBlock(latestBlock)
	if latestBlock > 0 {
		client.QueryBlock(latestBlock - 1)
	}

	// 2. 发送交易演示
	fmt.Println("2. 发送交易演示:")

	// 获取发送方地址和余额
	publicKey := getPublicKey(privateKey)
	senderAddress := crypto.PubkeyToAddress(*publicKey)
	balance, err := client.GetBalance(senderAddress.Hex())
	if err != nil {
		log.Fatalf("查询余额失败: %v", err)
	}

	fmt.Printf("发送方地址: %s\n", senderAddress.Hex())
	fmt.Printf("账户余额: %s ETH\n", weiToEther(balance))

	// 检查余额是否足够
	minBalance := big.NewInt(1000000000000000) // 0.001 ETH
	if balance.Cmp(minBalance) < 0 {
		fmt.Println("余额不足，请先从水龙头获取测试币")
		return
	}

	// 发送小额交易
	toAddress := "0x742d35Cc6634C0532925a3b8Dc2388e46F6Ee6b7" // 示例地址
	amount := big.NewInt(100000000000000)                     // 0.0001 ETH

	fmt.Printf("\n准备发送交易:\n")
	fmt.Printf("发送金额: %s ETH\n", weiToEther(amount))
	fmt.Printf("接收地址: %s\n", toAddress)

	// 取消注释以下代码来实际发送交易
	/*
		txHash, err := client.SendTransaction(privateKey, toAddress, amount)
		if err != nil {
			log.Fatalf("发送交易失败: %v", err)
		}
		fmt.Printf("交易发送成功! 交易哈希: %s\n", txHash)
	*/
}

// 辅助函数
func getPublicKey(privateKeyHex string) *ecdsa.PublicKey {
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)
	return privateKey.Public().(*ecdsa.PublicKey)
}

func weiToEther(wei *big.Int) string {
	ether := new(big.Float).SetInt(wei)
	ether = ether.Quo(ether, big.NewFloat(1e18))
	return ether.Text('f', 8) // 显示8位小数
}

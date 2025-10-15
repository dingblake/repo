package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/text"
)

func pollTransactionStatusWithDetail(signatureStr string) {
	signature, err := solana.SignatureFromBase58(signatureStr)
	if err != nil {
		panic(err)
	}

	rpcClient := rpc.New(rpc.DevNet_RPC)

	fmt.Printf("开始轮询交易 %s 的状态...\n", signatureStr)

	for i := 0; i < 30; i++ {
		// 获取交易详情
		tx, err := rpcClient.GetTransaction(
			context.TODO(),
			signature,
			&rpc.GetTransactionOpts{
				Encoding: solana.EncodingBase64,
			},
		)

		if err != nil {
			fmt.Printf("查询交易错误: %v\n", err)
		} else if tx != nil {
			fmt.Printf("✅ 交易找到! 区块: %d\n", tx.Slot)

			if tx.Meta != nil {
				if tx.Meta.Err == nil {
					fmt.Println("🎉 交易成功!")
					fmt.Printf("手续费: %d lamports\n", tx.Meta.Fee)
					fmt.Printf("前余额: %v\n", tx.Meta.PreBalances)
					fmt.Printf("后余额: %v\n", tx.Meta.PostBalances)
					return
				} else {
					fmt.Printf("❌ 交易失败: %v\n", tx.Meta.Err)
					return
				}
			}
		} else {
			fmt.Printf("⏳ 交易尚未确认... (%d/30)\n", i+1)
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("⏰ 轮询超时，交易可能尚未被网络确认")
}
func main() {
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	//查询余额
	pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
	out, err := client.GetBalance(
		context.TODO(),
		pubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}

	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

	var lamportsOnAccount = new(big.Float).SetUint64(uint64(out.Value))
	// Convert lamports to sol:
	var solBalance = new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

	// WARNING: this is not a precise conversion.
	fmt.Println("余额为", solBalance.Text('f', 10))
	//查询区块
	example, err := client.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	{
		maxVersion := uint64(0)
		out, err := client.GetBlockWithOpts(
			context.TODO(),
			uint64(example.Context.Slot),
			&rpc.GetBlockOpts{
				Encoding:                       solana.EncodingBase64,
				Commitment:                     rpc.CommitmentFinalized,
				TransactionDetails:             rpc.TransactionDetailsFull, // 获取完整交易信息
				MaxSupportedTransactionVersion: &maxVersion,                // 关键：添加这个参数
			},
		)
		if err != nil {
			panic(err)
		}
		// spew.Dump(out) // NOTE: This generates a lot of output.
		spew.Dump(len(out.Transactions))
	}

	{
		includeRewards := false
		maxVersion := uint64(0)
		out, err := client.GetBlockWithOpts(
			context.TODO(),
			uint64(example.Context.Slot),
			// You can specify more options here:
			&rpc.GetBlockOpts{
				Encoding:   solana.EncodingBase64,
				Commitment: rpc.CommitmentFinalized,
				// Get only signatures:
				TransactionDetails: rpc.TransactionDetailsSignatures,
				// Exclude rewards:
				Rewards:                        &includeRewards,
				MaxSupportedTransactionVersion: &maxVersion,
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}

	//GetBlockCommitment

	out1, err := client.GetBlockCommitment(
		context.TODO(),
		uint64(example.Context.Slot),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out1)

	//GetBlockProduction

	{
		out, err := client.GetBlockProduction(context.TODO())
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
	{
		out, err := client.GetBlockProductionWithOpts(
			context.TODO(),
			&rpc.GetBlockProductionOpts{
				Commitment: rpc.CommitmentFinalized,
				// Range: &rpc.SlotRangeRequest{
				//  FirstSlot: XXXXXX,
				//  Identity:  solana.MustPublicKeyFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				// },
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
	//	GetBlockTime
	out2, err := client.GetBlockTime(
		context.TODO(),
		uint64(example.Context.Slot),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out2)
	spew.Dump(out2.Time().Format(time.RFC1123))

	//GetHealth
	out3, err := client.GetHealth(
		context.TODO(),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out3)
	spew.Dump(out3 == rpc.HealthOk)

	//转账交易

	// Create a new RPC client:
	rpcClient := rpc.New(rpc.DevNet_RPC)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}
	defer wsClient.Close() // 记得关闭连接

	// Load the account that you will send funds FROM:
	//accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile("/home/ding/.config/solana/id.json")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	privateKeyPath := filepath.Join(homeDir, ".config", "solana", "id.json")
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile(privateKeyPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("accountFrom private key:", accountFrom)
	fmt.Println("accountFrom public key:", accountFrom.PublicKey())
	fmt.Println("=== 详细余额检查 ===")
	balance, err := rpcClient.GetBalance(context.TODO(), accountFrom.PublicKey(), rpc.CommitmentFinalized)
	if err != nil {
		fmt.Printf("获取余额错误: %v\n", err)
		panic(err)
	}
	fmt.Printf("账户地址: %s\n", accountFrom.PublicKey())
	fmt.Printf("余额: %d lamports (%.9f SOL)\n", balance.Value, float64(balance.Value)/float64(solana.LAMPORTS_PER_SOL))

	// 检查账户信息
	accInfo, err := rpcClient.GetAccountInfo(context.TODO(), accountFrom.PublicKey())
	if err != nil {
		fmt.Printf("获取账户信息错误: %v\n", err)
	} else {
		fmt.Printf("账户信息: %+v\n", accInfo)
	}
	// The public key of the account that you will send sol TO:
	accountTo := accountFrom.PublicKey()
	fmt.Printf("接收地址: %s\n", accountTo)
	// The amount to send (in lamports);
	// 1 sol = 1000000000 lamports
	amount := uint64(3333)

	// 直接检查余额是否足够，不再请求空投
	if balance.Value < amount {
		fmt.Printf("余额不足，需要 %d lamports，当前只有 %d lamports\n", amount, balance.Value)
		fmt.Println("请通过命令行请求空投: solana airdrop 2")
		return
	}
	//---------------

	recent, err := rpcClient.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amount,
				accountFrom.PublicKey(),
				accountTo,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(accountFrom.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if accountFrom.PublicKey().Equals(key) {
				return &accountFrom
			}
			return nil
		},
	)
	if err != nil {
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}
	spew.Dump(tx)
	// Pretty print the transaction:
	tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	// Send transaction, and wait for confirmation:
	sig, err := rpcClient.SendTransaction(context.TODO(), tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction signature:", sig)
	// 使用新的轮询函数获取详细交易信息
	fmt.Println("\n=== 获取交易详细信息 ===")
	pollTransactionStatusWithDetail(sig.String())
	// 等待交易确认
	fmt.Println("Waiting for transaction confirmation...")
	for i := 0; i < 30; i++ { // 最多等待60秒
		// 检查交易状态
		resp, err := rpcClient.GetSignatureStatuses(context.TODO(), false, sig)
		if err != nil {
			panic(err)
		}

		if resp != nil && len(resp.Value) > 0 && resp.Value[0] != nil {
			status := resp.Value[0]
			if status.ConfirmationStatus != "" {
				fmt.Printf("Confirmation status: %s\n", status.ConfirmationStatus)
				if status.ConfirmationStatus == rpc.ConfirmationStatusFinalized {
					fmt.Println("Transaction finalized!")
					break
				}
			}
		}

		time.Sleep(2 * time.Second)
	}
	spew.Dump(sig)

}

package handles

import (
	"context"
	"ethereum/config"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

// AddressVerify 验证地址
func AddressVerify(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

// ErrorFatal 重大错误 终止程序
func ErrorFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// ErrorPrint 可忽略错误 打印错误
func ErrorPrint(err error) {
	if err != nil {
		log.Println(err)
	}
}

// GetTransactionCost 获取交易gas费用
func GetTransactionCost(address string) uint64 {
	client := config.GetClient()
	txHash := common.HexToHash(address)
	receipt, err := client.TransactionReceipt(context.TODO(), txHash)
	if err != nil {
		log.Println(err)
		return 0
	}
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil || isPending {
		log.Println(err)
		return 0
	}
	return tx.GasPrice().Uint64() * receipt.GasUsed
}

package router

import (
	"ethereum/controllers/account"
	"ethereum/controllers/transaction"
	"ethereum/middleware"

	"github.com/gin-gonic/gin"
)

// GetRouter 返回所有路由
func GetRouter() *gin.Engine {
	app := gin.Default()

	v1 := app.Group("v1")
	{
		// 使用RSA中间件
		v1.Use(middleware.Rsa())
		// 根据地址查看用户账户
		v1.POST("balance", account.Balance)
		// ETH 转账
		v1.POST("ETHTransfer", account.ETHTransfer)
		// 提币
		v1.POST("extract", transaction.NewExtract)
	}
	// 生成账户
	app.POST("/v1/accountGenerate", account.Generate)
	// 获取一个用户
	app.POST("/v1/createAccount", account.NewAccount)
	// 代币转账
	app.POST("tokenTransfer", account.TokenTransfer)
	// ETH 转账
	app.POST("ETHTransfer", account.ETHTransfer)
	// 获取所有充值
	app.POST("getRechargeAddress", account.GetRechargeAddress)
	return app
}

package handles

import (
	"encoding/json"
	"ethereum/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 返回数据类型结构体
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ReturnJSON 返回json
func (rd ResponseData) ReturnJSON(c *gin.Context) {
	c.AbortWithStatusJSON(rd.Code, rd)
}

// Error 错误返回
func Error(code int, message string, c *gin.Context) {
	ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	}.ReturnJSON(c)
}

// Success 正确返回
func Success(message string, data interface{}, c *gin.Context) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	// 加密
	encryptData, err := config.RsaEncrypt(jsonData)
	if err != nil {
		log.Println(err)
	}
	ResponseData{
		Code:    http.StatusOK,
		Message: message,
		Data:    encryptData,
	}.ReturnJSON(c)
}

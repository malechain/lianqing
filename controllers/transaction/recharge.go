package transaction

import (
	"encoding/json"
	"ethereum/config"
	"ethereum/handles"
	"ethereum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewExtract 申请一个新的提币
func NewExtract(c *gin.Context) {
	var extract *models.Extract
	data, _ := c.Get("data")
	if err := json.Unmarshal(data.([]byte), &extract); err != nil {
		handles.Error(http.StatusBadRequest, "数据格式不正确", c)
		return
	}
	if !handles.AddressVerify(extract.From) || !handles.AddressVerify(extract.To) {
		handles.Error(http.StatusBadRequest, "提币地址不正确", c)
		return
	}
	// 从总账户给用户提供外部账户转账
	txhash, err := handles.TokenTransfer(extract.From, extract.To, extract.Value)
	if err != nil {
		handles.Error(http.StatusBadRequest, err.Error(), c)
		return
	}

	extract.TxHash = txhash
	extract.Create()
	rc := config.GetRedis()
	err = rc.SAdd("extract", txhash).Err()
	rc.Close()
	handles.Success("提交成功!", gin.H{"data": txhash}, c)
}

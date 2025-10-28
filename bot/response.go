package bot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 服务器错误
func serverError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	})
}

// 操作成功
func sendSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}

// 发送代码
func sendText(c *gin.Context, code int, text string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  text,
	})
}

// 发送数据
func sendData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

// 发送表格
func sendRows(c *gin.Context, total int64, data any) {
	type Data struct {
		Total int64 `json:"total,string"`
		Rows  any   `json:"rows"`
	}
	sendData(c, &Data{
		Total: total,
		Rows:  data,
	})
}

// 发送对象
func sendObject(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// serialize.go is a Go file that contains the code to serialize data to JSON format.

package core

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Brandon-lz/gmqtt/utils"

	"github.com/gin-gonic/gin"
)

type ApiOKResponse struct {
	Code    int    `json:"code"`    // 200
	Message string `json:"message"` // "success"
	Data    string `json:"data"`
}

// SuccessHandler 返回成功响应
func SuccessHandler(c *gin.Context, responseData interface{}) {
	if responseData == nil {
		c.JSON(http.StatusOK, ApiOKResponse{
			Code:    http.StatusOK,
			Message: "success",
			Data:    "",
		})
	} else {
		c.JSON(http.StatusOK, responseData)
	}
}

func SerializeDataAndValidate[T interface{}](source T, target *T, doSerialize ...bool) T { // target必须为指针类型
	if len(doSerialize) > 0 && doSerialize[0] {
		return SerializeData(source, target)
	} else {

		return ValidateSchema(source)
	}
}

func SerializeData[T interface{}](source any, target *T) T { // target必须为指针类型
	return utils.DeserializeData(source, target)
}

func ValidateSchema[T interface{}](source T) T {
	if err := ValidateStruct(source); err != nil {
		slog.Error(fmt.Sprintf("数据校验失败: %v", err))
		panic(NewKnownError(http.StatusInternalServerError, nil, "output数据异常"))
	}
	return source
}

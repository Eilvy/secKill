package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	param = Resp{
		Status:  20000,
		Message: "param error",
	}

	success = Resp{
		Status:  10000,
		Message: "success",
	}

	received = Resp{
		Status:  10001,
		Message: "received",
	}

	internal = Resp{
		Status:  20001,
		Message: "internal error",
	}

	parseErr = Resp{
		Status:  50000,
		Message: "parse token error",
	}
)

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, success)
}

func OKWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10000,
		"message": "success",
		"data":    data,
	})
}

func ParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, param)
}

func InternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, internal)
}

func ReceivedErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, received)
}

func ParseErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, parseErr)
}

package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc { // HandlerFunc is the type expected for all the middlewares
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d\n",
			param.ClientIP,
			param.TimeStamp,
			param.Method,
			param.Path,
			param.StatusCode,
		)
	})
}

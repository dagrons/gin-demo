package jwt

import (
	"net/http"
	"time"

	"github.com/dagrons/gin-demo/search_codimd/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var jwtSecret string

type option func()

func WithViperConfig() option {
	return func() {
		jwtSecret = viper.GetString("jwt_secret")
	}
}

func Option(opts ...option) {
	for _, opt := range opts {
		opt()
	}
}

func JWT(opts ...option) gin.HandlerFunc {
	Option(opts...)
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := ParseToken(token, jwtSecret)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

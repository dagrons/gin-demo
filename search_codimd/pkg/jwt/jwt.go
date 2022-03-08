package jwt

import (
	"net/http"
	"time"

	"github.com/dagrons/gin-demo/search_codimd/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type JwtClient struct {
	jwtSecret string
}

type option func(j *JwtClient)

func WithViperConfig() option {
	return func(j *JwtClient) {
		j.jwtSecret = viper.GetString("jwt_secret")
	}
}

func (j *JwtClient) Option(opts ...option) {
	for _, opt := range opts {
		opt(j)
	}
}

func New(opts ...option) *JwtClient {
	j := &JwtClient{}
	j.Option(opts...)
	return j
}

func JWT(opts ...option) gin.HandlerFunc {
	j := New(opts...)
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := ParseToken(token, j.jwtSecret)
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

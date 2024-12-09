package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (n *Network) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Bearer token을 가져온다.
		t := getAuthToken(c)

		if t == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort() // 나머지 핸들러들이 호출되지 않도록 하는 것. (대기중인 핸들러 호출 방지)
		} else {
			if _, err := n.gRPCClient.VerifyAuth(t); err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}

func getAuthToken(c *gin.Context) string {
	var token string

	authToken := c.Request.Header.Get("Authorization")
	// Bearer ~~~~~
	authSided := strings.Split(authToken, " ") // Bearer 문구 제거 후 실질적인 토큰 값 추출
	if len(authSided) > 1 {
		token = authSided[1]
	}

	return token
}

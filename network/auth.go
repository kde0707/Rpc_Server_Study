package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc-server/types"
)

func (n *Network) login(c *gin.Context) {
	// Auth Data 생성 필요
	var req types.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil { // 들어오는 파라미터를 parsing 할 수 있음.
		c.JSON(http.StatusBadRequest, err.Error())
	} else if res, err := n.service.CreateAuth(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (n *Network) verify(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}

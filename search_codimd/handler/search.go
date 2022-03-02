package handler

import (
	"github.com/dagrons/gin-demo/search_codimd/dal"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context, words []string) (interface{}, error) {
	return dal.Search(c, words)
}

package views

import (
	"log"

	"github.com/dagrons/gin-demo/search_codimd/handler"
	"github.com/dagrons/gin-demo/search_codimd/pkg/e"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	word_list := c.QueryArray("word_list")
	resultMap, err := handler.Search(c, word_list)
	if err != nil {
		log.Fatalf("fail to search, err=%v", err)
		c.JSON(e.ERROR, e.GetMsg(e.ERROR))
	}
	res := map[string]interface{}{
		"result": resultMap,
	}
	if err != nil {
		log.Fatalf("marshal failed, err=%v", err)
		c.JSON(e.ERROR, e.GetMsg(e.ERROR))
	}
	c.JSON(e.SUCCESS, res)
}

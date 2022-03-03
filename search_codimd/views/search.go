package views

import (
	"github.com/dagrons/gin-demo/search_codimd/handler"
	"github.com/dagrons/gin-demo/search_codimd/pkg/e"
	"github.com/dagrons/gin-demo/search_codimd/pkg/logging"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	word_list := c.QueryArray("word_list")
	resultMap, err := handler.Search(c, word_list)
	if err != nil {
		if e.IsErrInvalidParam(err) {
			logging.Warn("failed to search, err=%v", err)
			c.JSON(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS))
		} else {
			logging.Error("failed to search, err=%v", err)
			c.JSON(e.ERROR, e.GetMsg(e.ERROR))
		}
	} else {
		res := map[string]interface{}{
			"result": resultMap,
		}
		if err != nil {
			logging.Error("marshal failed, err=%v", err)
			c.JSON(e.ERROR, e.GetMsg(e.ERROR))
		}
		c.JSON(e.SUCCESS, res)
	}
}

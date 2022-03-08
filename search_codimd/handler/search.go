package handler

import (
	"fmt"

	"github.com/dagrons/gin-demo/search_codimd/dal"
	"github.com/dagrons/gin-demo/search_codimd/pkg/e"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context, words []string) (interface{}, error) {
	if len(words) == 0 || len(words[0]) == 0 {
		return nil, &e.ErrInvalidParam{
			Param: map[string]interface{}{
				"words": words,
			},
		}
	}
	resultMap, err := dal.Search(c, words)
	if err != nil {
		return nil, fmt.Errorf("call dal Search failed, err=%w", err)
	}
	return resultMap, nil
}

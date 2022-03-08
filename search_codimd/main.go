package main

import (
	"fmt"
	"net/http"

	"github.com/dagrons/gin-demo/search_codimd/dal"
	accesslog "github.com/dagrons/gin-demo/search_codimd/pkg/access_log"
	"github.com/dagrons/gin-demo/search_codimd/pkg/logging"
	_ "github.com/dagrons/gin-demo/search_codimd/pkg/settings"
	"github.com/dagrons/gin-demo/search_codimd/views"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	defer dal.Pg.Close()

	// 配置插件
	dal.Init(dal.WithViperConfig())
	logging.Init(logging.WithViperConfig())

	router := gin.Default()

	// 配置中间件
	router.Use(accesslog.Logger(accesslog.WithViperConfig()))

	// 配置路由
	router.GET("/api/search", views.Search)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("http_port")),
		Handler: router,
	}
	server.ListenAndServe()
}

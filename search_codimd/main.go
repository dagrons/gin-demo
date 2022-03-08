package main

import (
	"fmt"
	"net/http"

	"github.com/dagrons/gin-demo/search_codimd/dal"
	"github.com/dagrons/gin-demo/search_codimd/pkg/jwt"
	"github.com/dagrons/gin-demo/search_codimd/pkg/logging"
	_ "github.com/dagrons/gin-demo/search_codimd/pkg/settings"
	"github.com/dagrons/gin-demo/search_codimd/views"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	logging.Init(logging.WithViperConfig()) // package本身就是个单例，这是个非常重要的思想
	jwt.Init(jwt.WithViperConfig())         // option function pattern, 也是个很优雅的设计，将闭包函数作为参数传递

	defer dal.Pg.Close()

	router := gin.Default()
	router.GET("/api/search", views.Search)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("http_port")),
		Handler: router,
	}
	server.ListenAndServe()
}

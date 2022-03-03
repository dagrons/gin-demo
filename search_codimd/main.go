package main

import (
	"fmt"
	"net/http"

	"github.com/dagrons/gin-demo/search_codimd/dal"
	"github.com/dagrons/gin-demo/search_codimd/views"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	defer dal.Pg.Close()

	router := gin.Default()
	router.GET("/api/search", views.Search)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("http_port")),
		Handler: router,
	}
	server.ListenAndServe()
}

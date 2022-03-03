package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dagrons/gin-demo/search_codimd/dal"
	"github.com/dagrons/gin-demo/search_codimd/pkg/utils"
	"github.com/dagrons/gin-demo/search_codimd/views"
	"github.com/gin-gonic/gin"
)

func main() {
	defer dal.Pg.Close()

	fmt.Println(os.Getenv("conf_dir"))

	router := gin.Default()
	router.GET("/api/search", views.Search)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", utils.GetEnvInt("http_port", 8089)),
		Handler: router,
	}
	server.ListenAndServe()
}

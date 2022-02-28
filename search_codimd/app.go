package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dagrons/gin-demo/search_codimd/dal"
	"github.com/dagrons/gin-demo/search_codimd/views"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

var (
	port int
)

func init() {
	serverCfg, err := ini.Load("conf/server.ini")
	if err != nil {
		panic(err)
	}
	port = serverCfg.Section("server").Key("port").MustInt()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	defer dal.Pg.Close()
	router := gin.Default()
	router.GET("/api/search", views.Search)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	server.ListenAndServe()
}

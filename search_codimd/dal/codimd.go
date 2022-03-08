package dal

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/dagrons/gin-demo/search_codimd/pkg/settings"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	Pg             *sql.DB
	pg_conf_string string
)

func Init(opts ...option) {
	var err error
	Option(opts...)
	Pg, err = sql.Open("postgres", pg_conf_string)
	if err != nil {
		panic("error connect to postgres")
	}
}

func WithViperConfig() option {
	return func() {
		pgCfg := viper.GetStringMapString("postgres")
		pg_conf_string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
			pgCfg["hostname"],
			pgCfg["username"],
			pgCfg["password"],
			pgCfg["dbname"],
			pgCfg["sslmode"])
	}
}

type option func()

func Option(opts ...option) {
	for _, opt := range opts {
		opt()
	}
}

func Search(c *gin.Context, words []string) (interface{}, error) {
	resultMap := make(map[string]string)
	var sqlStr = `SELECT title, shortid FROM "Notes"`
	prunedWords := make([]string, 0)
	for _, word := range words { // avoid sql injection
		var prunedWord strings.Builder
		for i := 0; i < len(word); i++ {
			if word[i] == '\'' {
				prunedWord.WriteString("''")
			} else {
				prunedWord.WriteString("'")
			}
		}
		prunedWords = append(prunedWords, word)
	}
	n := len(prunedWords)
	if n > 0 {
		sqlStr += fmt.Sprintf(` WHERE (title LIKE '%%%s%%' OR content LIKE '%%%s%%')`, prunedWords[0], prunedWords[0])
	}
	if n > 1 {
		for _, word := range words[1:] {
			sqlStr += fmt.Sprintf(` AND (title LIKE '%%%s%%' OR content LIKE '%%%s%%')`, word, word)
		}
	}
	rows, err := Pg.Query(sqlStr)
	if err != nil {
		return nil, fmt.Errorf("pg query failed, err=%w, sqlStr=%s", err, sqlStr)
	}
	for rows.Next() {
		var (
			title   string
			shortId string
		)
		if err := rows.Scan(&title, &shortId); err != nil { // 反序列化
			return resultMap, err
		}
		resultMap[title] = shortId
	}
	return resultMap, nil
}

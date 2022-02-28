package dal

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
)

var Pg *sql.DB

func init() {
	var err error
	cfg, err := ini.Load("conf/db.ini")
	if err != nil {
		panic(err)
	}
	pgSection := cfg.Section("postgres")
	pg_conf_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		pgSection.Key("hostname"),
		pgSection.Key("username"),
		pgSection.Key("password"),
		pgSection.Key("dbname"),
		pgSection.Key("sslmode"))
	Pg, err = sql.Open("postgres", pg_conf_string)
	if err != nil {
		panic(err)
	}
}

func Search(c *gin.Context, words []string) ([]interface{}, error) {
	resultSet := []interface{}{}
	var sqlStr = `SELECT title FROM "Notes"`
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
		var elem interface{}
		if err := rows.Scan(&elem); err != nil { // 反序列化
			return resultSet, err
		}
		resultSet = append(resultSet, elem)
	}
	return resultSet, nil
}

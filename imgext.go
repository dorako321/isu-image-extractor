package main

import (
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"flag"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"./modules/Binary"
)

var (
	tbName     = flag.String("tb", "image", "table name (default is image)")
	pkColumn   = flag.String("pc", "id", "primary key column name (default is id)")
	dataColumn = flag.String("dc", "data", "binary data column name (default is data)")

	host      = flag.String("h", "localhost", "host name (default is localhost)")
	port      = flag.String("p", "3306", "port number (default is 3306)")
	user      = flag.String("u", "root", "user name (default is root)")
	pass      = flag.String("pw", "password", "password (default is password)")
	dbName    = flag.String("d", "isucon", "database name (default is isucon)")
	outputDir = flag.String("o", "./output", "output directory (default is ./output)")
)

func main() {
	// パラメーターの取得
	flag.Parse()
	// パスワードの取得
	if *pass == "" {
		fmt.Print("Password: ")
		pw, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		} else {
			*pass = string(pw)
		}
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		*user,
		*pass,
		*host,
		*port,
		*dbName,
	)

	// connect mysql db
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()

	// select
	// 必要に応じてクエリを修正してください
	rows, err := db.Query(
		"SELECT `" + *pkColumn + "`, `" + *dataColumn +
			"` FROM `" + *tbName + "`")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	var id string
	var data []byte
	for rows.Next() {
		err = rows.Scan(&id, &data)

		// バイナリデータから拡張子を判定
		ext := Binary.GetExtensionName(data)
		fmt.Println(id, ext)
		ioutil.WriteFile(*outputDir + "/" + id + ext, data, 0666)
		os.Exit(0)
	}
}
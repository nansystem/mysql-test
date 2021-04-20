package main

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"github.com/nansystem/mysql-test/usecase"
)

var opts struct {
	TargetDB string `short:"t" long:"targetdb" choice:"cvad" required:"true"`
	Count    int64  `short:"c" long:"count" default:"10"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usecase.FillCvs(db, opts.Count)
}

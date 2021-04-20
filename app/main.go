package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nansystem/mysql-test/generated"
)

var opts struct {
	TargetDB string `short:"t" long:"targetdb" choice:"cvad" required:"true"`
}

func main() {
	// args, err := flags.Parse(&opts)
	// if err != nil {
	// 	os.Exit(1)
	// }
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ad := &generated.Cv{
		AdID:      1,
		UserID:    2,
		Status:    3,
		CreatedAt: 4,
	}
	ad.Insert(db)
}

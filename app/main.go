package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nansystem/mysql-test/usecase"
)

// import (
// 	"database/sql"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jessevdk/go-flags"
// 	"github.com/nansystem/mysql-test/usecase"
// )

// var opts struct {
// 	TargetDB string `short:"t" long:"targetdb" choice:"cvad" required:"true"`
// 	Count    int64  `short:"c" long:"count" default:"10"`
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	_, err := flags.Parse(&opts)
// 	if err != nil {
// 		os.Exit(1)
// 	}

// 	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/mysql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	usecase.FillCvs(db, opts.Count)
// }

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id      int64  `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     int64  `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/devdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usecase.FillCvs(db, 728332)
	// clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	// if err != nil {
	// 	panic(err)
	// }
	// defer clientsFile.Close()

	// clients := []*Client{}
	// for i := 0; i < 100000; i++ {
	// 	clients = append(clients, &Client{Id: utils.RandNum(1, 1000), Name: "John Jhones", Age: utils.RandNum(16, 100)}) // Add clients
	// }
	// err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	// if err != nil {
	// 	panic(err)
	// }

}

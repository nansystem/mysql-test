package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nansystem/mysql-test/infra"
	usecase "github.com/nansystem/mysql-test/usecase/order_master"
)

func main() {
	err := infra.InitDB("root:password@tcp(db:3306)/devdb")
	if err != nil {
		log.Fatal(err)
	}
	defer infra.Close()

	usecase.FillOrderMaster(1000000)
}

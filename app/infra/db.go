package infra

import "database/sql"

var DB *sql.DB

func InitDB(dataSourceName string) error {
	var err error

	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	return DB.Ping()
}

func Close() {
	DB.Close()
}

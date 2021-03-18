package main

import (
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Company struct {
	id            int
	name          string
	category      string
	address       string
	url           string
	employees_num int
}

func main() {
	// CSVを読み込む
	file, err := os.Open("./test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	var line []string

	// 最初の2行を無視する
	line, err = reader.Read()
	line, err = reader.Read()

	// DBに接続
	db, err := sql.Open("mysql", "root:@tcp(mysql:3306)/test_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 行数分のLoop
	var index int
	for {
		index = index + 1

		line, err = reader.Read()
		if err != nil {
			break
		}

		// DB書き込み
		stmtInsert, err := db.Prepare(
			"INSERT INTO companies(id, name, category, address, url, employees_num) VALUES(?, ?, ?, ?, ?, ?)",
		)
		if err != nil {
			panic(err.Error())
		}
		defer stmtInsert.Close()

		if line[12] == "-" {
			line[12] = ""
		}

		var company Company
		company.id = index
		company.name = line[0]
		company.category = line[16]
		company.address = line[24]
		company.url = line[3]
		company.employees_num, _ = strconv.Atoi(line[12])

		_, err = stmtInsert.Exec(
			company.id,
			company.name,
			company.category,
			company.address,
			company.url,
			company.employees_num,
		)
		if err != nil {
			panic(err.Error())
		}
	}
}

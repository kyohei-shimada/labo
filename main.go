package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/sqldef/schema"
)

func main() {
	desiredSQL := `CREATE TABLE sample (
		name varchar(255) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`

	currentSQL := `CREATE TABLE sample (
		name varchar(255) DEFAULT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`

	diffs, err := schema.GenerateIdempotentDDLs(schema.GeneratorModeMysql, desiredSQL, currentSQL)
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", diffs)
}

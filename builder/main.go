package main

import (
	querybuilder "designPatterns/builder/queryBuilder"
	"fmt"
)

func main() {

	clientCode(querybuilder.NewMysqlQueryBuilder())

	clientCode(querybuilder.NewPostgresQueryBuilder())
}

func clientCode(builder querybuilder.SQLQueryBuilder) {
	builder.Select("users", []string{"id", "name"}).
		Where("id", "1", "=").
		Limit(0, 10)
	fmt.Println(builder.GetSQL())
}

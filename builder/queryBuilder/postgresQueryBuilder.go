package querybuilder

import "fmt"

type PostgresQueryBuilder struct {
	MysqlQueryBuilder // Встраивание, а не наследование
}

func NewPostgresQueryBuilder() *PostgresQueryBuilder {
	return &PostgresQueryBuilder{MysqlQueryBuilder: *NewMysqlQueryBuilder()}
}

func (b *PostgresQueryBuilder) Limit(start, offset int) SQLQueryBuilder {
	b.query.limit = fmt.Sprintf(" LIMIT %d OFFSET %d", start, offset)
	return b
}

package querybuilder

import (
	"fmt"
	"strings"
)

type MysqlQueryBuilder struct {
	query *queryBuilder
}

type queryBuilder struct {
	base  string
	typ   string
	where []string
	limit string
}

func NewMysqlQueryBuilder() *MysqlQueryBuilder {
	return &MysqlQueryBuilder{}
}

func (m *MysqlQueryBuilder) reset() {
	m.query = &queryBuilder{}

}

// Select начинает построение запроса SELECT.
func (m *MysqlQueryBuilder) Select(table string, fields []string) SQLQueryBuilder {
	m.reset()
	m.query.base = "SELECT " + strings.Join(fields, ", ") + " FROM " + table
	m.query.typ = "select"
	return m
}

// Where добавляет условие WHERE к запросу.
func (m *MysqlQueryBuilder) Where(field, value, operator string) SQLQueryBuilder {
	if m.query.typ != "select" && m.query.typ != "update" && m.query.typ != "delete" {
		panic("WHERE can only be added to SELECT, UPDATE, OR DELETE")
	}
	m.query.where = append(m.query.where, fmt.Sprintf("%s %s '%s'", field, operator, value))
	return m
}

func (m *MysqlQueryBuilder) Limit(start, offset int) SQLQueryBuilder {
	if m.query.typ != "select" {
		panic("LIMIT can only be added to SELECT")
	}
	m.query.limit = fmt.Sprintf(" LIMIT %d, %d", start, offset)
	return m
}

func (m *MysqlQueryBuilder) GetSQL() string {
	sql := m.query.base
	if len(m.query.where) > 0 {
		sql += " WHERE " + strings.Join(m.query.where, " AND ")
	}
	if m.query.limit != "" {
		sql += m.query.limit
	}
	sql += ";"
	return sql
}

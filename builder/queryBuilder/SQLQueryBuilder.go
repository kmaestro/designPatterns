package querybuilder

type SQLQueryBuilder interface {
	Select(table string, fields []string) SQLQueryBuilder
    Where(field, value, operator string) SQLQueryBuilder
    Limit(start, offset int) SQLQueryBuilder
    GetSQL() string
}

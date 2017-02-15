package db

import (
	"github.com/gocraft/dbr"
)

type Builder struct {
	d dbr.Dialect
}

func (this *Builder) Build(stmt dbr.Builder) (string, error) {
	return dbr.InterpolateForDialect("?", []interface{}{stmt}, this.d)
}

func (this *Builder) MustBuild(stmt dbr.Builder) string {
	query, err := this.Build(stmt)
	if err != nil {
		panic(err)
	}

	return query
}

func (Builder) InsertInto(table string) *dbr.InsertStmt {
	return dbr.InsertInto(table)
}

func (Builder) InsertBySQL(query string, value ...interface{}) *dbr.InsertStmt {
	return dbr.InsertBySql(query, value...)
}

func (this *Builder) Select(colunm ...string) *dbr.SelectStmt {
	c := make([]interface{}, len(colunm))
	for i, one := range colunm {
		c[i] = one
	}

	return this.SelectEx(c...)
}

func (Builder) SelectEx(colunm ...interface{}) *dbr.SelectStmt {
	return dbr.Select(colunm...)
}

func (Builder) SelectBySQL(query string, value ...interface{}) *dbr.SelectStmt {
	return dbr.SelectBySql(query, value...)
}

func (Builder) Update(table string) *dbr.UpdateStmt {
	return dbr.Update(table)
}

func (Builder) UpdateBySQL(query string, value ...interface{}) *dbr.UpdateStmt {
	return dbr.UpdateBySql(query, value...)
}

func (Builder) DeleteFrom(table string) *dbr.DeleteStmt {
	return dbr.DeleteFrom(table)
}

func (Builder) DeleteBySQL(query string, value ...interface{}) *dbr.DeleteStmt {
	return dbr.DeleteBySql(query, value...)
}

func (Builder) Expr(query string, value ...interface{}) dbr.Builder {
	return dbr.Expr(query, value...)
}

func (Builder) And(cond ...dbr.Builder) dbr.Builder {
	return dbr.And(cond...)
}

func (Builder) Or(cond ...dbr.Builder) dbr.Builder {
	return dbr.Or(cond...)
}

func (Builder) Eq(column string, value interface{}) dbr.Builder {
	return dbr.Eq(column, value)
}

func (Builder) Neq(column string, value interface{}) dbr.Builder {
	return dbr.Neq(column, value)
}

func (Builder) Gt(column string, value interface{}) dbr.Builder {
	return dbr.Gt(column, value)
}

func (Builder) Gte(column string, value interface{}) dbr.Builder {
	return dbr.Gte(column, value)
}

func (Builder) Lt(column string, value interface{}) dbr.Builder {
	return dbr.Lt(column, value)
}

func (Builder) Lte(column string, value interface{}) dbr.Builder {
	return dbr.Lte(column, value)
}

func (Builder) I(s string) dbr.I {
	return dbr.I(s)
}

func (Builder) Where(where interface{}, conds ...interface{}) Where {
	return Where{
		Where: where,
		Conds: conds,
	}
}

func (Builder) Order(col string, desc bool) Order {
	return Order{
		Column: col,
		DESC:   desc,
	}
}

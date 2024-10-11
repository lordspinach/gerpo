package sql

import (
	"context"
	"strconv"

	"github.com/insei/gerpo/types"
)

type StringSelectBuilder struct {
	ctx     context.Context
	columns []types.Column
	limit   uint64
	offset  uint64
	orderBy string
}

func (b *StringSelectBuilder) Columns(cols ...types.Column) {
	for _, col := range cols {
		b.columns = append(b.columns, col)
	}
}

func (b *StringSelectBuilder) Limit(limit uint64) {
	b.limit = limit
}

func (b *StringSelectBuilder) Offset(offset uint64) {
	b.offset = offset
}

func (b *StringSelectBuilder) OrderBy(columnDirection string) *StringSelectBuilder {
	if b.orderBy != "" {
		b.orderBy += ", "
	}
	b.orderBy += columnDirection
	return b
}

func (b *StringSelectBuilder) OrderByColumn(col types.Column, direction types.OrderDirection) error {
	if col.IsAllowedAction(types.SQLActionSort) {
		if b.orderBy != "" {
			b.orderBy += ", "
		}
		b.orderBy += col.ToSQL(b.ctx) + " " + string(direction)
	}
	return nil
}

func (b *StringSelectBuilder) GetColumns() []types.Column {
	return b.columns
}

func (b *StringSelectBuilder) GetSQL() string {
	sql := ""
	for _, col := range b.columns {
		if sql != "" {
			sql += ", "
		}
		sql += col.ToSQL(b.ctx)
	}
	return sql
}

func (b *StringSelectBuilder) GetOrderSQL() string {
	return b.orderBy
}

func (b *StringSelectBuilder) GetLimit() string {
	if b.limit == 0 {
		return ""
	}
	return strconv.FormatUint(b.limit, 10)
}

func (b *StringSelectBuilder) GetOffset() string {
	if b.offset == 0 {
		return ""
	}
	return strconv.FormatUint(b.offset, 10)
}

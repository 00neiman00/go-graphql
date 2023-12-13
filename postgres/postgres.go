package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type DbLogger struct{}

func (d DbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	fmt.Println(string(query))
	return nil
}

func New(options *pg.Options) *pg.DB {
	return pg.Connect(options)
}

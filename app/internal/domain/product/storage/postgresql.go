package storage

import (
	"GoService/app/internal/domain/product/model"
	"GoService/app/pkg/client/postgresql"
	db "GoService/app/pkg/client/postgresql/model"
	"GoService/app/pkg/logging"
	"context"
	"github.com/Masterminds/squirrel"
)

type ProductStorage struct {
	queryBuilder squirrel.StatementBuilderType
	client       PostgreSQLClient
	logger       *logging.Logger
}

func NewProductStorage(client PostgreSQLClient, logger *logging.Logger) ProductStorage {
	return ProductStorage{
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client:       client,
		logger:       logger,
	}
}

const (
	scheme = "public"
	table  = "product"
)

func (s *ProductStorage) queryLogger(sql, table string, args []interface{}) *logging.Logger {
	return s.logger.ExtraFields(map[string]interface{}{
		"sql":   sql,
		"table": table,
		"args":  args,
	})
}

func (s *ProductStorage) All(ctx context.Context) ([]model.Product, error) {
	query := s.queryBuilder.Select("id").
		Column("name").
		Column("description").
		Column("image_id").
		Column("price").
		Column("currency_id").
		Column("rating").
		Column("category_id").
		Column("specification").
		Column("created_at").
		Column("updated_at").
		From(scheme + "." + table)

	sql, args, err := query.ToSql()
	logger := s.queryLogger(sql, table, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err)
		return nil, err
	}

	logger.Trace("do query")
	rows, err := s.client.Query(ctx, sql, args...)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {
		p := model.Product{}
		if err = rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.ImageId,
			&p.Price,
			&p.CurrencyId,
			&p.Rating,
			&p.CategoryId,
			&p.Specification,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			err = db.ErrScan(postgresql.ParsePgError(err))
			logger.Error(err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

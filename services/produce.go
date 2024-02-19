package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/j7nw4r/produce-store/schemas"
	"log/slog"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
)

const SelectProduceSql = "select * from main.produce where id = ?"
const SelectProduceByLikeNameSql = "select * from main.produce where name like ? || '%'"
const SelectProduceByLikeCodeSql = "select * from main.produce where code like ? || '%'"
const InsertProduceSqlReturning = "insert into produce (code, name, price)  values (?, ?, ?) returning *"
const DeleteProduceSqlReturning = "delete from produce where id = ? returning *"

type ProduceService struct {
	db *sql.DB
}

func NewProduceService(db *sql.DB) ProduceService {
	return ProduceService{db: db}
}

func (ps ProduceService) GetProduce(ctx context.Context, id int) (*schemas.ProduceSchema, error) {
	if id <= 0 {
		return nil, ErrBadRequest
	}

	stmt, err := ps.db.PrepareContext(ctx, SelectProduceSql)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("could seletect for id: %s", id)
	}

	row := stmt.QueryRowContext(ctx, id)
	if row == nil {
		return nil, errors.New("could not query db")
	}

	prod := schemas.ProduceSchema{}
	if err := row.Scan(&prod.Id, &prod.Code, &prod.Name, &prod.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.New("could not map db row to produce")
	}

	return &prod, nil
}

func (ps ProduceService) DeleteProduce(ctx context.Context, id int) (*schemas.ProduceSchema, error) {
	if id <= 0 {
		return nil, ErrBadRequest
	}

	stmt, err := ps.db.PrepareContext(ctx, DeleteProduceSqlReturning)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("could not delete for id: %s", id)
	}

	row := stmt.QueryRowContext(ctx, id)
	if row == nil {
		return nil, errors.New("could not query db")
	}

	prod := schemas.ProduceSchema{}
	if err := row.Scan(&prod.Id, &prod.Code, &prod.Name, &prod.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.New("could not map db row to produce")
	}

	return &prod, nil
}

func (ps ProduceService) SearchProduceByName(ctx context.Context, name string) ([]schemas.ProduceSchema, error) {
	stmt, err := ps.db.PrepareContext(ctx, SelectProduceByLikeNameSql)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("could not search for %s", name)
	}

	rows, err := stmt.QueryContext(ctx, name)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("could not search for %s", name)
	}

	prods := []schemas.ProduceSchema{}
	for rows.Next() {
		var prod schemas.ProduceSchema
		if err := rows.Scan(&prod.Id, &prod.Code, &prod.Name, &prod.Price); err != nil {
			return nil, fmt.Errorf("could not search for %s", name)
		}

		prods = append(prods, prod)

	}

	return prods, nil
}

func (ps ProduceService) SearchProduceByCode(ctx context.Context, code string) ([]schemas.ProduceSchema, error) {
	stmt, err := ps.db.PrepareContext(ctx, SelectProduceByLikeCodeSql)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("could not search for code: %s", code)
	}

	rows, err := stmt.QueryContext(ctx, code)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("could not search for code: %s", code)
	}

	prods := []schemas.ProduceSchema{}
	for rows.Next() {
		var prod schemas.ProduceSchema
		if err := rows.Scan(&prod.Id, &prod.Code, &prod.Name, &prod.Price); err != nil {
			return nil, fmt.Errorf("could not search for code: %s", code)
		}
		prods = append(prods, prod)
	}

	return prods, nil
}

func (ps ProduceService) StoreProduce(ctx context.Context, pp []schemas.ProduceSchema) ([]schemas.ProduceSchema, error) {
	txn, err := ps.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("could not insert produce: %v", pp)
	}

	stmt, err := txn.PrepareContext(ctx, InsertProduceSqlReturning)
	if err != nil {
		txn.Rollback()
		slog.Error(err.Error())
		return nil, fmt.Errorf("could not insert produce: %v", pp)
	}

	retProds := []schemas.ProduceSchema{}
	for _, p := range pp {
		row := stmt.QueryRowContext(ctx, p.Code, p.Name, p.Price)
		if err != nil {
			txn.Rollback()
			slog.Error(err.Error())
			return nil, fmt.Errorf("could not insert produce: %v", p)
		}
		prod := schemas.ProduceSchema{}
		if err := row.Scan(&prod.Id, &prod.Code, &prod.Name, &prod.Price); err != nil {
			txn.Rollback()
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrNotFound
			}
			return nil, errors.New("could not map db row to produce")
		}
		retProds = append(retProds, prod)
	}

	txn.Commit()
	return retProds, nil
}

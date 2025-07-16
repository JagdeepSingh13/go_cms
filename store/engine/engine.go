package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/JagdeepSingh13/go_cms/models"
	"github.com/google/uuid"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) EngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("trnasaction rollback error")
			}
		} else {
			if cmErr := tx.Commit; cmErr != nil {
				fmt.Printf("trnasaction commit error")
			}
		}
	}()

	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id=$1", id).Scan(
		&engine.EngineID,
		&engine.Displacement,
		&engine.NoOfCylenders,
		&engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}

	return engine, nil
}

func (e EngineStore) EngineCreate(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("trnasaction rollback error")
			}
		} else {
			if cmErr := tx.Commit; cmErr != nil {
				fmt.Printf("trnasaction commit error")
			}
		}
	}()

	engineId := uuid.New()

	_, err = tx.ExecContext(ctx, "INSERT INTO engine (id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4)", engineId, engineReq.Displacement, engineReq.NoOfCylenders, engineReq.CarRange)
	if err != nil {
		return models.Engine{}, err
	}

	engine := models.Engine{
		EngineID:      engineId,
		Displacement:  engineReq.Displacement,
		NoOfCylenders: engineReq.NoOfCylenders,
		CarRange:      engineReq.CarRange,
	}
	return engine, nil
}

func (e EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) EngineDelete(ctx context.Context, id string) (models.Engine, error) {

}

package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Debasish-87/featurepilot/internal/organization/service"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

var _ service.Repository = (*PostgresRepository)(nil)

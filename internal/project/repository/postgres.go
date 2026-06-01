package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	projectService "github.com/Debasish-87/featurepilot/internal/project/service"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

var _ projectService.Repository = (*PostgresRepository)(nil)

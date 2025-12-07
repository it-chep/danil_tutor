package dal

import (
	"context"

	"github.com/it-chep/danil_tutor.git/internal/module/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) SetState(ctx context.Context, studentID int64, state dto.State) error {
	sql := `
		update students set state = $2 where id = $1
	`

	_, err := r.pool.Exec(ctx, sql, studentID, state)
	return err
}

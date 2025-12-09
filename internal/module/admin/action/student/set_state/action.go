package set_state

import (
	"context"
	"fmt"
	"github.com/it-chep/danil_tutor.git/internal/module/admin/dto"

	"github.com/it-chep/danil_tutor.git/internal/module/admin/action/student/set_state/dal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Repository
}

func New(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, studentID int64, state dto.State) error {
	if !state.Valid() {
		return fmt.Errorf("invalid state")
	}
	return a.dal.SetState(ctx, studentID, state)
}

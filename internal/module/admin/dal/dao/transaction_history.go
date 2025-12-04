package dao

import (
	"github.com/it-chep/danil_tutor.git/internal/module/admin/dto"
	"github.com/it-chep/danil_tutor.git/internal/pkg/convert"
	"github.com/it-chep/danil_tutor.git/pkg/xo"
)

type TransactionHistoryDAO struct {
	xo.TransactionsHistory
}

type TransactionsHistoryDAO []TransactionHistoryDAO

func (th *TransactionHistoryDAO) ToDomain() dto.TransactionHistory {
	return dto.TransactionHistory{
		ID:          th.ID,
		OrderID:     th.OrderID.String,
		CreatedAt:   th.CreatedAt,
		IsConfirmed: th.ConfirmedAt.Valid,
		Amount:      convert.NumericToDecimal(th.Amount),
		StudentID:   th.StudentID,
	}
}

func (ths *TransactionsHistoryDAO) ToDomain() []dto.TransactionHistory {
	domain := make([]dto.TransactionHistory, 0, len(*ths))
	for _, th := range *ths {
		domain = append(domain, th.ToDomain())
	}
	return domain
}

package get_all_finance

import (
	"context"
	indto "github.com/it-chep/danil_tutor.git/internal/module/admin/dto"
	"github.com/samber/lo"
	"sync"

	"github.com/it-chep/danil_tutor.git/internal/module/admin/action/get_all_finance/dal"
	"github.com/it-chep/danil_tutor.git/internal/module/admin/action/get_all_finance/dto"
	"github.com/it-chep/danil_tutor.git/internal/pkg/convert"
	"github.com/it-chep/danil_tutor.git/internal/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type Action struct {
	dal *dal.Repository
}

func New(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, from, to string, adminID int64) (dto.GetAllFinanceDto, error) {
	fromTime, toTime, err := convert.StringsIntervalToTime(from, to)
	if err != nil {
		return dto.GetAllFinanceDto{}, err
	}

	var (
		cashFlow   decimal.Decimal
		finance    decimal.Decimal
		conversion float64
		wg         = sync.WaitGroup{}
	)

	// Получаем общий оборот
	wg.Add(1)
	go func() {
		defer wg.Done()
		gCashFlow, gErr := a.dal.GetCashFlow(ctx, fromTime, toTime, adminID)
		if gErr != nil {
			logger.Error(ctx, "Ошибка при получении оборота", gErr)
			return
		}
		cashFlow = gCashFlow
	}()

	// Получаем расходы на зарплаты
	wg.Add(1)
	go func() {
		defer wg.Done()
		gfinance, gErr := a.dal.GetFinanceInfo(ctx, fromTime, toTime, adminID)
		if gErr != nil {
			logger.Error(ctx, "Ошибка при расходов на зп", gErr)
			return
		}
		finance = gfinance
	}()

	// Получаем конверсию
	wg.Add(1)
	go func() {
		defer wg.Done()
		gConversion, gErr := a.getConversion(ctx, adminID)
		if gErr != nil {
			logger.Error(ctx, "Ошибка при расходов на зп", gErr)
			return
		}
		conversion = gConversion
	}()

	wg.Wait()

	return dto.GetAllFinanceDto{
		Profit:     finance.String(),
		CashFlow:   cashFlow.String(),
		Conversion: conversion,
	}, nil
}

func (a *Action) getConversion(ctx context.Context, adminID int64) (float64, error) {
	students, err := a.dal.GetAllStudents(ctx, adminID)
	if err != nil {
		return 0, err
	}

	statusesMap := make(map[indto.State]int64)
	lo.ForEach(students, func(student indto.Student, _ int) {
		statusesMap[student.State]++
	})

	// Числитель
	numerator := int64(len(students))
	// Знаменатель
	denominator := int64(len(students)) -
		statusesMap[indto.NEW] -
		statusesMap[indto.IN_PROGRESS] -
		statusesMap[indto.DECLINED_AFTER_TRIAL] -
		statusesMap[indto.DECLINED_AFTER_LESSONS]

	if denominator == 0 {
		return 0, nil
	}

	//Конверсия = Рабочие чаты / (Всего чатов - Без пробного - В процессе - Отказ после пробного - Отказ после занятий)
	conversion := decimal.NewFromInt(numerator).
		Div(decimal.NewFromInt(denominator)).
		Mul(decimal.NewFromInt(100)).
		InexactFloat64()

	return conversion, nil
}

package bot

import (
	"context"
	"fmt"
	"github.com/it-chep/danil_tutor.git/internal/pkg/logger"
	"strconv"
	"strings"

	"github.com/it-chep/danil_tutor.git/internal/module/bot/action/commands/start"
	"github.com/it-chep/danil_tutor.git/internal/module/bot/dto"
)

func (b *Bot) Route(ctx context.Context, msg dto.Message) error {
	if strings.Contains(msg.Text, "id_") {
		atoi, err := strconv.Atoi(msg.Text[len("id_"):])
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("ошибка при парсинге ID: %s", msg.Text), err)
			return err
		}
		return b.Actions.AuthUser.Do(ctx, msg, int64(atoi))
	}

	switch msg.Text {
	case "/start":
		return b.Actions.Start.Start(ctx, msg)
	case "/info":
		return b.Actions.Info.Do(ctx, msg)
	case start.GetBalance, "/balance":
		return b.Actions.GetBalance.GetBalance(ctx, msg)
	case start.TopUpBalance, "/add_balance":
		return b.Actions.TopUpBalance.InitTransaction(ctx, msg)
	case start.GetLessons, "/lessons":
		return b.Actions.GetLessons.GetLessons(ctx, msg)
	default:
		if b.Actions.TopUpBalance.TransactionExists(ctx, msg) {
			return b.Actions.TopUpBalance.SetAmount(ctx, msg)
		}
	}
	return nil
}

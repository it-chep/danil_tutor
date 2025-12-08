package dto

import (
	"github.com/it-chep/danil_tutor.git/internal/module/admin/dto"
)

type FilterRequest struct {
	IsLost      bool
	TgUsernames []string
	States      dto.States
}

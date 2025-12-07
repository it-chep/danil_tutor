package filter_students

import (
	"github.com/it-chep/danil_tutor.git/internal/module/admin/action/student/filter_students/dto"
	common "github.com/it-chep/danil_tutor.git/internal/module/dto"
	"github.com/samber/lo"
)

type Request struct {
	AdminsUsernames []string `json:"tg_admins_usernames"`
	IsLost          bool     `json:"is_lost"`
	States          []int    `json:"states,omitempty"`
}

func (r Request) ToFilterRequest() dto.FilterRequest {
	return dto.FilterRequest{
		IsLost:      r.IsLost,
		TgUsernames: r.AdminsUsernames,
		States: lo.Map(r.States, func(s int, _ int) common.State {
			return common.State(s)
		}),
	}
}

package get_states

import (
	"encoding/json"
	"net/http"

	"github.com/it-chep/danil_tutor.git/internal/module/dto"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := h.prepareResponse()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) prepareResponse() Response {
	return Response{
		States: []State{
			{
				State: dto.NEW.Int(),
				Name:  "Новый",
				Desc:  "Без пробного занятия",
			},
			{
				State: dto.WORKING.Int(),
				Name:  "Рабочий",
				Desc:  "Просто посещает занятия",
			},
			{
				State: dto.IN_PROGRESS.Int(),
				Name:  "В процессе",
				Desc:  "Договариваются о пробном занятии",
			},
			{
				State: dto.DECLINED_AFTER_TRIAL.Int(),
				Name:  "Отказ",
				Desc:  "Отказ после пробного занятия",
			},
			{
				State: dto.BEFORE_TRIAL.Int(),
				Name:  "Не дошел",
				Desc:  "Не дошел до пробного занятия",
			},
			{
				State: dto.DECLINED_AFTER_LESSONS.Int(),
				Name:  "Отказ после занятий",
				Desc:  "Отказ после занятий",
			},
		},
	}
}

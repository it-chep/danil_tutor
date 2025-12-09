package set_state

import (
	"encoding/json"
	"github.com/it-chep/danil_tutor.git/internal/module/admin/dto"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/it-chep/danil_tutor.git/internal/module/admin"
)

type Handler struct {
	adminModule *admin.Module
}

func NewHandler(adminModule *admin.Module) *Handler {
	return &Handler{
		adminModule: adminModule,
	}
}

func (h *Handler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		studentIDStr := chi.URLParam(r, "student_id")
		studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid user ID", http.StatusBadRequest)
			return
		}

		err = h.adminModule.Actions.SetState.Do(ctx, studentID, dto.State(req.State))
		if err != nil {
			http.Error(w, "failed to get user data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

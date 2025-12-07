package dal

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/danil_tutor.git/internal/module/admin/dal/dao"
	"github.com/it-chep/danil_tutor.git/internal/module/admin/dto"
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

func (r *Repository) GetStudent(ctx context.Context, studentID int64) (dto.Student, error) {
	sql := `
		select * from students where id = $1
	`
	var student dao.StudentDAO
	err := pgxscan.Get(ctx, r.pool, &student, sql, studentID)
	if err != nil {
		return dto.Student{}, err
	}

	return student.ToDomain(), nil
}

func (r *Repository) GetTutor(ctx context.Context, tutorID int64) (dto.Tutor, error) {
	sql := `
		select 
		    t.id,
			t.cost_per_hour,
			t.subject_id,
			t.admin_id,
			u.full_name,
			u.tg,
			u.phone 
		from tutors t 
		    join users u on t.id = u.tutor_id 
		where t.id = $1
	`

	args := []interface{}{
		tutorID,
	}

	var tutor dao.TutorDAO
	err := pgxscan.Get(ctx, r.pool, &tutor, sql, args...)
	if err != nil {
		return dto.Tutor{}, err
	}

	return tutor.ToDomain(), nil
}

func (r *Repository) GetStudentWallet(ctx context.Context, studentID int64) (dto.Wallet, error) {
	sql := `
		select * from wallet where student_id = $1
	`
	var wallet dao.Wallet
	err := pgxscan.Get(ctx, r.pool, &wallet, sql, studentID)
	if err != nil {
		return dto.Wallet{}, err
	}
	return wallet.ToDomain(), nil
}

func (r *Repository) UpdateStudentWallet(ctx context.Context, studentID int64, remain decimal.Decimal) error {
	sql := `
		update wallet set balance = $1 where student_id = $2
	`

	args := []interface{}{
		remain,
		studentID,
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}

// ConductLesson помечаем что урок проведен
func (r *Repository) ConductLesson(ctx context.Context, tutorID, studentID, durationInMinutes int64, createdTime time.Time) error {
	sql := `
		insert into conducted_lessons(student_id, tutor_id, duration_in_minutes, is_trial, created_at, is_first_paid_lesson)
		values ($1, $2, $3, false, $4, false)
	`

	args := []interface{}{
		studentID,
		tutorID,
		durationInMinutes,
		createdTime.UTC(),
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}

// ConductFirstPaidLesson помечаем что урок проведен
func (r *Repository) ConductFirstPaidLesson(ctx context.Context, tutorID, studentID, durationInMinutes int64, createdTime time.Time) error {
	sql := `
		insert into conducted_lessons(student_id, tutor_id, duration_in_minutes, is_trial, created_at, is_first_paid_lesson)
		values ($1, $2, $3, false, $4, true)
	`

	args := []interface{}{
		studentID,
		tutorID,
		durationInMinutes,
		createdTime.UTC(),
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}

// FinishTrial помечаем что урок проведен
func (r *Repository) FinishTrial(ctx context.Context, studentID int64) error {
	sql := `
		update students set is_finished_trial = true where id = $1
	`
	_, err := r.pool.Exec(ctx, sql, studentID)
	return err
}

// HasFirstPaidLesson .
func (r *Repository) HasFirstPaidLesson(ctx context.Context, tutorID, studentID int64) (bool, error) {
	sql := `	
		select * from conducted_lessons where tutor_id = $1 and student_id = $2 and is_first_paid_lesson = true	
	`

	args := []interface{}{
		tutorID,
		studentID,
	}

	var lessons dao.ConductedLessonDAOs
	err := pgxscan.Select(ctx, r.pool, &lessons, sql, args...)
	if err != nil {
		return true, err
	}

	if len(lessons) != 0 {
		return true, nil
	}

	return false, nil
}

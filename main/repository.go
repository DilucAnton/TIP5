package main

import (
	"context"
	"database/sql"
	"time"
)

// Task — модель для сканирования результатов SELECT
type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo { return &Repo{DB: db} }

// CreateTask — параметризованный INSERT с возвратом id
// Изменила чтобы можно было указать done
func (r *Repo) CreateTask(ctx context.Context, title string, done bool) (int, error) {
	var id int
	const q = `INSERT INTO tasks (title, done) VALUES ($1,  $2) RETURNING id;`
	err := r.DB.QueryRowContext(ctx, q, title, done).Scan(&id)
	return id, err
}

// ListTasks — базовый SELECT всех задач (демо для занятия)
func (r *Repo) ListTasks(ctx context.Context) ([]Task, error) {
	const q = `SELECT id, title, done, created_at FROM tasks ORDER BY id;`
	rows, err := r.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *Repo) ListDone(ctx context.Context, done bool) ([]Task, error) {
	const q = `SELECT id, title, done, created_at FROM tasks WHERE done = $1 ORDER BY id;`

	rows, err := r.DB.QueryContext(ctx, q, done)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// FindByID поиск по ID задачи
func (r *Repo) FindByID(ctx context.Context, id int) (*Task, error) {
	const q = `SELECT id, title, done, created_at FROM tasks WHERE id = $1;`

	var task Task
	err := r.DB.QueryRowContext(ctx, q, id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)
	// err nil
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &task, nil
}

// массовая вставка
func (r *Repo) CreateMany(ctx context.Context, tasks []struct {
	title string
	done  bool
}) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		_, err := tx.ExecContext(ctx,
			"INSERT INTO tasks (title, done) VALUES ($1, $2)",
			task.title, task.done)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

package memberships

import (
	"context"
	"database/sql"

	"github.com/tamaaa13/fastcampus/internal/model/memberships"
)

func (r *repository) GetUSer(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error) {
	query := "SELECT id, email, username, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ? OR id = ?"
	row := r.db.QueryRowContext(ctx, query, email, username, userID)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Username, &response.Password, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBY, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := "INSERT INTO users (email, username, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Username, model.Password, model.CreatedAt, model.UpdatedAt, model.CreatedBY, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

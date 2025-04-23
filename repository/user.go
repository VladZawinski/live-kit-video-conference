package repository

import (
	"database/sql"
	"errors"

	"github.com/live-kit-video-conference/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int) (*model.User, error)
	GetAll() ([]*model.User, error)
	Delete(id int) error
}

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) Create(user *model.User) error {
	query := `INSERT INTO user (username) VALUES (?, ?)`
	result, err := r.DB.Exec(query, user.Username)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

func (r *userRepositoryImpl) GetByID(id int) (*model.User, error) {
	query := `SELECT id, username, created_at FROM user WHERE id = ?`
	row := r.DB.QueryRow(query, id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) GetAll() ([]*model.User, error) {
	query := `SELECT id, username, created_at FROM user`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepositoryImpl) Delete(id int) error {
	query := `DELETE FROM user WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

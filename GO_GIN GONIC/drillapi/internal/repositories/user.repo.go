package repositories

import (
	"julianmindria/drill/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `INSERT INTO public.users (
        name
    )
    VALUES(
        :name
    )`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "user has been added", nil
}

func (r RepoUser) GetUser(data *models.User) ([]models.User, error) {
	users_data := []models.User{}
	if err := r.Select(&users_data, `select * FROM public.users`); err != nil {
		return nil, err
	}

	return users_data, nil
}

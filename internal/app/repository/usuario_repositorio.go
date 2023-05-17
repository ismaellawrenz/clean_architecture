package repository

import (
	"crud-golang/internal/app/model"
	"database/sql"

	"github.com/pkg/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {

	stmt, err := r.db.Prepare("INSERT INTO usuario (nome, email, senha) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare statement")
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return user, nil
}

func (r *UserRepository) GetByID(id int64) (*model.User, error) {

	stmt, err := r.db.Prepare("SELECT id, nome, email FROM usuario WHERE id = $1")
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare statement")
	}
	defer stmt.Close()

	user := &model.User{}
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return user, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {

	stmt, err := r.db.Prepare("UPDATE public.usuario	SET nome=$1, email=$2, senha=$3	WHERE id = $4;")
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Query(user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return user, nil
}

func (r *UserRepository) Delete(id int64) error {

	stmt, err := r.db.Prepare("DELETE FROM public.usuario	WHERE id = $1")
	if err != nil {
		return errors.Wrap(err, "failed to prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return errors.Wrap(err, "failed to execute query")
	}

	return nil
}

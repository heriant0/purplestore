package repositories

import (
	"errors"
	"fmt"

	"github.com/heriant0/purplestore/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Register(data models.User) error {
	sqlStatement := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
	`

	_, err := r.DB.Exec(sqlStatement, data.Email, data.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var (
		data         models.User
		sqlStatement = `
			SELECT id, email, password
			FROM users
			WHERE email = $1
			LIMIT 1 
		`
	)
	fmt.Println(sqlStatement, email)
	err := r.DB.QueryRowx(sqlStatement, email).StructScan(&data)
	if err != nil {
		return nil, err
	}

	if data.ID == 0 {
		return nil, errors.New("data not found")
	}

	return &data, nil
}

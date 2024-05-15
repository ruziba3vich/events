package postgres

import (
	"database/sql"
	"errors"

	"github.com/ruziba3vich/events/models"
	auth "github.com/ruziba3vich/authentication_tokens"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) LogIn() error {
	return nil
}

func (u *UserRepo) CreateEvent() error {
	return nil
}

func (u *UserRepo) Register(req models.RegisterRequest) (string, error) {
	if hashedPwd, err := hashPassword(req.Password); err != nil {
		return "", errors.New("error while hashing password")
	} else {
		req.Password = hashedPwd
	}
	query := `
		INSERT INTO Users (
			username,
			password
		)
		VALUES (
			$1, $2
		)
		RETURNING id, username, password;
	`
	var user models.User
	row, err := u.db.Query(query, req.Username, req.Password)
	if err != nil {
		return "", errors.New("unable to add the user into the database")
	}
	defer row.Close()
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.Id, req.Username)

	if err != nil {
		return "", errors.New("error while generating token")
	}
	return token, err
}

func (u *UserRepo) UpdateEvent() error {
	return nil
}

func (u *UserRepo) DeleteEvent() error {
	return nil
}

func (u *UserRepo) GetAllEvents() error {
	return nil
}

/*
Register()
	LogIn()
	CreateEvent()
	UpdateEvent()
	DeleteEvent()
	GetAllEvents()
*/

/*
if hashedPwd, err := hashPassword(req.Password); err != nil {
		return nil, errors.New("error while hashing password")
	} else {
		req.Password = hashedPwd
	}
	query := `
		INSERT INTO Users (
			username,
			password
		)
		VALUES (
			$1, $2
		)
		RETURNING id, username, password;
	`
	var user models.User
	row, err := s.db.Query(query, req.Username, req.Password)
	if err != nil {
		return nil, errors.New("unable to add the user into the database")
	}
	defer row.Close()
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
*/

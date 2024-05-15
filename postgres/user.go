package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	auth "github.com/ruziba3vich/authentication_tokens"
	"github.com/ruziba3vich/events/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) LogIn(req models.RegisterRequest) (string, error) {
	if hashPassword, err := hashPassword(req.Password); err != nil {
		return "", errors.New("error while hashing password")
	} else {
		req.Password = hashPassword
	}
	query := `
		SELECT id, username FROM Users
		WHERE username = $1 AND password = $2;
	`
	var user models.User
	row := u.db.QueryRow(query, req.Username, req.Password)
	if err := row.Scan(&user.Id, &user.Username); err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.Id, req.Username)

	if err != nil {
		return "", errors.New("error while creating token")
	}
	return token, nil
}

func (u *UserRepo) CreateEvent(req models.CreateEventRequest) (*models.Event, error) {
	query := `INSERT INTO Events (
		user_id,
		title,
		description,
		event_time
	) VALUES (
		$1, $2, $3, $4
	) RETURNING id, user_id, title, description, event_time;
	`
	var event models.Event

	row := u.db.QueryRow(query,
		req.UserId,
		req.Title,
		req.Description,
		req.EventTime)
	if err := row.Scan(&event.Id, &event.UserId, &event.Title, &event.Description, &event.EventTime); err != nil {
		return nil, err
	}
	query = `
		INSERT INTO Files (
			file_name,
			event_id
		) VALUES (
			$1, $2
		) RETURNING file_name;
	`

	for i := range req.Files {
		row := u.db.QueryRow(query, req.Files[i], event.Id)
		var fileName string
		if err := row.Scan(&fileName); err != nil {
			return nil, err
		}
		event.Files = append(event.Files, fileName)
	}

	return &event, nil
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

func (u *UserRepo) UpdateEvent(req models.UpdateEventRequest) (*EventDTO, error) {
	query := `
	UPDATE Events e
	SET e.title = $1,
		e.description = $2,
		e.event_time = $3
	WHERE e.id = $4
	RETURING e.id, e.title, e.description, e.event_time;
	`
	var event EventDTO
	row := u.db.QueryRow(query, req.Title, req.Description, req.EventTime, req.Id)
	if err := row.Scan(&event.Id, &event.Title, &event.Description, &event.EventTime); err != nil {
		return nil, err
	}
	query = `
		SELECT file_name FROM Files WHERE user_id = $1;
	`
	rows, err := u.db.Query(query, event.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fileName string
		if err := rows.Scan(&fileName); err != nil {
			return nil, err
		}
		event.Files = append(event.Files, fileName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &event, nil
}

func (u *UserRepo) DeleteEvent(req models.DeleteEventRequest) error {
	exists, err := u.checkEventExists(req.EventId)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("event does not exist")
	}

	err = u.deleteEvent(req.EventId)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) checkEventExists(eventID uuid.UUID) (bool, error) {
	var exists bool
	err := u.db.QueryRow("SELECT EXISTS(SELECT 1 FROM Events WHERE id = $1)", eventID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u *UserRepo) deleteEvent(eventID uuid.UUID) error {
	_, err := u.db.Exec("DELETE FROM Events WHERE id = $1", eventID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetAllEvents(req models.GetAllEventsRequest) ([]EventDTO, error) {
	query := `
		SELECT e.id, e.title, e.description, e.event_time
		FROM Events e WHERE e.user_id = $1;
	`
	rows, err := u.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []EventDTO

	for rows.Next() {
		var eventDTO EventDTO
		err := rows.Scan(&eventDTO.Id, &eventDTO.Title, &eventDTO.Description, &eventDTO.EventTime)
		if err != nil {
			return nil, err
		}
		newQuery := `
			SELECT file_name WHERE event_id = $1;
		`
		newRows, e := u.db.Query(newQuery, eventDTO.Id)
		if e != nil {
			return nil, e
		}
		defer newRows.Close()
		for newRows.Next() {
			var fileName string
			if e := newRows.Scan(&fileName); e != nil {
				return nil, e
			}
			eventDTO.Files = append(eventDTO.Files, fileName)
		}
		if e := newRows.Err(); e != nil {
			return nil, e
		}
		events = append(events, eventDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

type EventDTO struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventTime   time.Time `json:"event_time"`
	Files       []string  `json:"files"`
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

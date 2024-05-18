package user

import (
	"database/sql"
	"fmt"

	"github.com/dev-by-sjb/e-commerce-go-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email  = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := rows.Scan(
		u.ID,
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) CreateUser(u *types.User) error {
	_, err := s.db.Exec("INSERT INTO users(first_name, last_name, email, hashed_password) VALUES($1 $2 $3 $4 $5)", u.FirstName, u.LastName, u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByID(id int64) (*types.User, error) {
	return nil, nil
}

package user

import (
	"errors"
	"time"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) IsValid() error {
	if u.Username == "" {
		return errors.New("Username é obrigatório")
	}

	if len(u.Username) < 3 {
		return errors.New("Username deve ter pelo menos 3 letras.")
	}

	if len(u.Username) > 50 {
		return errors.New("username deve ter no máximo 50 letras")
	}

	if u.Password == "" {
		return errors.New("Password é obrigatório.")
	}

	return nil
}

func (u *User) HasValidUsername() bool {
	return len(u.Username) >= 3 && len(u.Username) <= 50
}

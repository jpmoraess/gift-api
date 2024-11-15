package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type User struct {
	id        uuid.UUID
	username  string
	password  string
	fullName  string
	email     string
	createdAt time.Time
}

func NewUser(username, password, fullName, email string) (user *User, err error) {
	user = &User{
		id:        uuid.New(),
		username:  username,
		password:  password,
		fullName:  fullName,
		email:     email,
		createdAt: time.Now(),
	}

	if err = user.validate(); err != nil {
		return
	}

	return
}

func RestoreUser(id uuid.UUID, username string, password string, fullName string, email string, createdAt time.Time) (user *User, err error) {
	user = &User{
		id:        id,
		username:  username,
		password:  password,
		fullName:  fullName,
		email:     email,
		createdAt: createdAt,
	}

	if err = user.validate(); err != nil {
		return
	}

	return
}

func (u *User) validate() error {
	if u.id == uuid.Nil {
		return errors.New("id is required")
	}

	if len(u.username) == 0 {
		return errors.New("username is required")
	}

	if len(u.password) == 0 {
		return errors.New("password is required")
	}

	if len(u.fullName) == 0 {
		return errors.New("fullName is required")
	}

	if len(u.email) == 0 {
		return errors.New("email is required")
	}

	return nil
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) FullName() string {
	return u.fullName
}

func (u *User) Email() string {
	return u.email
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

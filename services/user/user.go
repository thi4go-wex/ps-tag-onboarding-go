package user

import (
	"ps-tag-onboarding-go/db"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type UserServiceI interface {
	UserGet(id string) (*User, error)
	UserSet(user *User) error
}

type UserService struct {
	db db.Database
}

func NewUserService(db db.Database) (*UserService, error) {
	return &UserService{
		db: db,
	}, nil
}

func (s *UserService) UserGet(id string) (*User, error) {
	data, err := s.db.Get(db.DBUserCollection, id)
	if err != nil {
		return nil, err
	}

	user := data.(User)

	return &user, nil
}

func (s *UserService) UserSet(user *User) error {
	return s.db.Set(db.DBUserCollection, user)
}

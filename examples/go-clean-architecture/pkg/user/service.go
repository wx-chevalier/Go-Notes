package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type Service interface {
	Register(ctx context.Context, email, password, phoneNumber string) (*User, error)

	Login(ctx context.Context, email, password string) (*User, error)

	ChangePassword(ctx context.Context, email, password string) error

	BuildProfile(ctx context.Context, user *User) (*User, error)

	GetUserProfile(ctx context.Context, email string) (*User, error)

	IsValid(user *User) (bool, error)

	GetRepo() Repository
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Register(ctx context.Context, email, password, phoneNumber string) (u *User, err error) {

	exists, err := s.repo.DoesEmailExist(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("User already exists")
	}

	hasher := md5.New()
	hasher.Write([]byte(password))

	return s.repo.CreateMinimal(ctx, email, hex.EncodeToString(hasher.Sum(nil)), phoneNumber)
}

func (s *service) Login(ctx context.Context, email, password string) (u *User, err error) {

	hasher := md5.New()
	hasher.Write([]byte(password))
	return s.repo.FindByEmailAndPassword(ctx, email, hex.EncodeToString(hasher.Sum(nil)))
}

func (s *service) ChangePassword(ctx context.Context, email, password string) (err error) {

	hasher := md5.New()
	hasher.Write([]byte(password))
	return s.repo.ChangePassword(ctx, email, hex.EncodeToString(hasher.Sum(nil)))
}

func (s *service) BuildProfile(ctx context.Context, user *User) (u *User, err error) {

	return s.repo.BuildProfile(ctx, user)
}

func (s *service) GetUserProfile(ctx context.Context, email string) (u *User, err error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *service) IsValid(user *User) (ok bool, err error) {

	return ok, err
}

func (s *service) GetRepo() Repository {

	return s.repo
}

package user

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id uint) (*User, error)

	BuildProfile(ctx context.Context, user *User) (*User, error)

	CreateMinimal(ctx context.Context, email, password, phoneNumber string) (*User, error)

	FindByEmailAndPassword(ctx context.Context, email, password string) (*User, error)

	FindByEmail(ctx context.Context, email string) (*User, error)

	DoesEmailExist(ctx context.Context, email string) (bool, error)

	ChangePassword(ctx context.Context, email, password string) error
}

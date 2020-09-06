package admin

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id uint) (*Admin, error)

	BuildProfile(ctx context.Context, user *Admin) (*Admin, error)

	CreateMinimal(ctx context.Context, email, password, phoneNumber string) (*Admin, error)

	FindByEmailAndPassword(ctx context.Context, email, password string) (*Admin, error)

	FindByEmail(ctx context.Context, email string) (*Admin, error)

	DoesEmailExist(ctx context.Context, email string) (bool, error)

	ChangePassword(ctx context.Context, email, password string) error
}

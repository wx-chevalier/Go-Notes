package user

import (
	"context"

	pkg "github.com/L04DB4L4NC3R/jobs-mhrd/pkg"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) FindByID(ctx context.Context, id uint) (user *User, err error) {

	return user, err
}

func (r *repo) BuildProfile(ctx context.Context, user *User) (u *User, err error) {

	result := r.DB.Table("users").Where("email = ?", user.Email).Updates(map[string]interface{}{
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"phone_number": user.PhoneNumber,
		"address":      user.Address,
		"display_pic":  user.DisplayPic,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	switch result.Error {
	case nil:
		return user, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) CreateMinimal(ctx context.Context, email, password, phoneNumber string) (u *User, err error) {

	u = &User{
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}
	result := r.DB.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (r *repo) FindByEmailAndPassword(ctx context.Context, email, password string) (u *User, err error) {

	u = &User{}
	result := r.DB.Where("email = ? AND password = ?", email, password).First(u)

	switch result.Error {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) DoesEmailExist(ctx context.Context, email string) (doesEmailExist bool, err error) {

	u := &User{}
	if r.DB.Where("email = ?", email).First(u).RecordNotFound() {
		return false, nil
	}

	return true, nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (u *User, err error) {

	u = &User{}
	projection := "email, created_at, updated_at, deleted_at, phone_number, first_name, last_name, address, display_pic"
	result := r.DB.Select(projection).Where("email = ?", email).First(u)

	switch result.Error {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) ChangePassword(ctx context.Context, email, password string) error {

	result := r.DB.Table("users").Where("email = ?", email).Update(map[string]interface{}{
		"password": password,
	})

	switch result.Error {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return pkg.ErrNotFound
	default:
		return pkg.ErrDatabase
	}
}

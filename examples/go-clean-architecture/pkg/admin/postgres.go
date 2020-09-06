package admin

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

func (r *repo) FindByID(ctx context.Context, id uint) (admin *Admin, err error) {

	return admin, err
}

func (r *repo) BuildProfile(ctx context.Context, admin *Admin) (u *Admin, err error) {

	result := r.DB.Table("admins").Where("email = ?", admin.Email).Updates(map[string]interface{}{
		"first_name":   admin.FirstName,
		"last_name":    admin.LastName,
		"phone_number": admin.PhoneNumber,
		"address":      admin.Address,
		"display_pic":  admin.DisplayPic,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	switch result.Error {
	case nil:
		return admin, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) CreateMinimal(ctx context.Context, email, password, phoneNumber string) (u *Admin, err error) {

	u = &Admin{
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

func (r *repo) FindByEmailAndPassword(ctx context.Context, email, password string) (u *Admin, err error) {

	u = &Admin{}
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

	u := &Admin{}
	if r.DB.Where("email = ?", email).First(u).RecordNotFound() {
		return false, nil
	}

	return true, nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (u *Admin, err error) {

	u = &Admin{}
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

	result := r.DB.Table("admins").Where("email = ?", email).Update(map[string]interface{}{
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

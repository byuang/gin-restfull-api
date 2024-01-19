package repository

import (
	"errors"
	"gin-restfull-api/data/request"
	"gin-restfull-api/helper"
	"gin-restfull-api/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (u *UsersRepositoryImpl) Delete(usersId int) {
	var users model.Users
	result := u.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

func (u *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	results := u.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}

func (u *UsersRepositoryImpl) FindById(usersId int) (model.Users, error) {
	var users model.Users
	result := u.Db.Find(&users, usersId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("users is not found")
	}
}

func (u *UsersRepositoryImpl) Save(users model.Users) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

func (u *UsersRepositoryImpl) Update(users model.Users) {
	var updateUsers = request.UpdateUsersRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}

func (u *UsersRepositoryImpl) FindByUsername(username string) (model.Users, error) {
	var users model.Users
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}

func (u *UsersRepositoryImpl) FindByEmail(email string) (model.Users, error) {
	var users model.Users
	result := u.Db.First(&users, "email = ?", email)

	if result.RowsAffected == 0 {
		return users, errors.New("email not found")
	}

	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (u *UsersRepositoryImpl) UpdateOtp(users model.Users) {
	updateFields := map[string]interface{}{
		"Password" : users.Password,
		"PasswordResetToken": users.PasswordResetToken,
		"PasswordResetAt":    users.PasswordResetAt,
	}

	result := u.Db.Model(&users).Updates(updateFields)
	helper.ErrorPanic(result.Error)
}

func (u *UsersRepositoryImpl) FindByOtp(Otp int) (model.Users, error) {
	var users model.Users
	result := u.Db.First(&users, "password_reset_token = ?", Otp)

	if result.RowsAffected == 0 {
		return users, errors.New("otp not found")
	}

	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

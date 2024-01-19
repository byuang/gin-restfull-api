package repository

import "gin-restfull-api/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (model.Users, error)
	FindAll() []model.Users
	FindByUsername(username string) (model.Users, error)
	FindByEmail(email string) (model.Users, error)
	UpdateOtp(users model.Users)
	FindByOtp(otp int) (model.Users, error)
}

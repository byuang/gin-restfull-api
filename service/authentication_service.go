package service

import "gin-restfull-api/data/request"

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUsersRequest)
}

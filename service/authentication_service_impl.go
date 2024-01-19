package service

import (
	"errors"
	"fmt"
	"gin-restfull-api/config"
	"gin-restfull-api/data/request"
	"gin-restfull-api/helper"
	"gin-restfull-api/model"
	"gin-restfull-api/repository"
	"gin-restfull-api/utils"
	"time"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
	}
}

//Login
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {	
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

//Register
func (a *AuthenticationServiceImpl) Register(users request.CreateUsersRequest) {

	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UsersRepository.Save(newUser)
}


// ForgotPassword
func (a *AuthenticationServiceImpl) ForgotPassword(users request.ForgotPasswordRequest) (string, error) {
    existingUser, err := a.UsersRepository.FindByEmail(users.Email)
    if err != nil {
        return "", errors.New("Email not found")
    }

    otp := utils.GenerateOTP(4)
    if err != nil {
        return "", errors.New("failed to generate token otp")
    }

    existingUser.PasswordResetToken = otp
    existingUser.PasswordResetAt = time.Now().Add(time.Minute * 5)

    a.UsersRepository.UpdateOtp(existingUser)

	emailData := utils.EmailData{
		Otp: otp,
		Email: existingUser.Email,
		Subject: " Reset Password",
	}

	utils.SendEmail(&existingUser, &emailData, "resetPassword.html")

    return fmt.Sprintf("%d", otp), nil
}


// ResetPassword
func (a *AuthenticationServiceImpl) ResetPassword(otp int, users request.ResetPasswordRequest) (string, error) {
	
	existingUser, err := a.UsersRepository.FindByOtp(otp)
	if err != nil {
		return "", errors.New("Otp not found")
	}

	if otp != existingUser.PasswordResetToken {
		return "", errors.New("Invalid OTP")
	}

	if time.Now().After(existingUser.PasswordResetAt) {
		return "", errors.New("OTP has expired")
	}

	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	existingUser.Password = hashedPassword
	existingUser.PasswordResetToken = 0
	existingUser.PasswordResetAt = time.Time{}

	a.UsersRepository.UpdateOtp(existingUser)
	return "", nil
}

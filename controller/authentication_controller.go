package controller

import (
	"fmt"
	"gin-restfull-api/data/request"
	"gin-restfull-api/data/response"
	"gin-restfull-api/helper"
	"gin-restfull-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
	Validate       *validator.Validate
}

func NewAuthenticationController(service service.AuthenticationService,  validate *validator.Validate) *AuthenticationController {
	return &AuthenticationController{
		authenticationService: service,
		Validate:       validate,
	}
	
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	
	if err := controller.Validate.Struct(loginRequest); err != nil {
		if helper.ValidationError(err, ctx, loginRequest) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "BadRequest",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	if err := controller.Validate.Struct(createUsersRequest); err != nil {
		if helper.ValidationError(err, ctx, createUsersRequest) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	controller.authenticationService.Register(createUsersRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) ForgotPassword(ctx *gin.Context) {

	forgotPasswordRequest := request.ForgotPasswordRequest{}
	err := ctx.ShouldBindJSON(&forgotPasswordRequest)
	helper.ErrorPanic(err)

	if err := controller.Validate.Struct(forgotPasswordRequest); err != nil {
		if helper.ValidationError(err, ctx, forgotPasswordRequest) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	otp, err := controller.authenticationService.ForgotPassword(forgotPasswordRequest)

	if err != nil {
        webResponse := response.Response{
            Code:    http.StatusInternalServerError,
            Status:  "Internal Server Error",
            Message: "Failed to process forgot password request",
            Data:    nil,
        }
        ctx.JSON(http.StatusInternalServerError, webResponse)
        return
    }

	otpStr, err := strconv.Atoi(otp)
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Forgot Passworsssd successfully",
		Data:   otpStr,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) ResetPassword(ctx *gin.Context) {

	otpStr := ctx.Query("otp")

    otp, err := strconv.Atoi(otpStr)
    helper.ErrorPanic(err)

    resetPasswordRequest := request.ResetPasswordRequest{}
    err = ctx.ShouldBindJSON(&resetPasswordRequest)
    helper.ErrorPanic(err)

	if err := controller.Validate.Struct(resetPasswordRequest); err != nil {
		if helper.ValidationError(err, ctx, resetPasswordRequest) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
    controller.authenticationService.ResetPassword(otp, resetPasswordRequest)

    webResponse := response.Response{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "Reset Password successfully",
        Data:    nil,
    }
    ctx.JSON(http.StatusOK, webResponse)
}


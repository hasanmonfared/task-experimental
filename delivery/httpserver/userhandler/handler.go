package userhandler

import (
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
)

type Handler struct {
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}

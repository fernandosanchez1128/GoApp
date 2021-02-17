package Middleware

import (
	"GoApp/Model"
	"context"
	"github.com/go-kit/kit/log"
)

type ErrorHandler struct {
	Logger log.Logger
}

func (errorHandler ErrorHandler) Handle(ctx context.Context, err error) {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			errorHandler.Logger.Log("error", err)
		}
	}()
	customErr := err.(Model.CustomError)
	errorHandler.Logger.Log(
		"error", customErr,
		"rootError", customErr.OrigError)

}

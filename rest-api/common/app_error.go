package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (a *AppError) RootError() error {
	if err, ok := a.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return a.RootErr
}

func (a *AppError) Error() string {
	return a.RootErr.Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrCantGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCantGet%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with server", err.Error(), "ErrInternal")
}

func ErrInvalidRequest(err error) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrEntityExisted(err error, entity string) *AppError {
	return NewFullErrorResponse(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%v already exist", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("Err%vExist", entity))
}

func ErrCannotCreateEntity(err error, entity string) *AppError {
	return NewFullErrorResponse(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("can not create %v", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("ErrCannotCreate%v", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		"dont have permission",
		"ErrNoPermission",
	)
}

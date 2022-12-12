package service

import (
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// InternalServerError ...
	InternalServerError = "Internal Server Error"

	// AlreadyExistsError ...
	AlreadyExistsError = "Already Exists"

	// NotFoundError ...
	NotFoundError = "Not Found"

	// InvalidArgumentError ...
	InvalidArgumentError = "Invalid Argument"
)

var (

	// ErrNotFound ...
	ErrNotFound = errors.New("not found")

	// ErrInternal ...
	ErrInternalServer = errors.New("internal server error")

	// ErrAlreadyExists ...
	ErrAlreadyExists = errors.New(" already exists")

	// ErrUsernameExists ...
	ErrUsernameExists = errors.New("username exists")

	//ErrPhoneExists
	ErrPhoneExists = errors.New("phone exists")

	// ErrEmailExists ...
	ErrEmailExists = errors.New("email exists")

	// ErrInvalidField ...
	ErrInvalidField = errors.New("invalid field for username/email")

	// ErrMaximumAmount ...
	ErrMaximumAmount = errors.New("maximum amount")

	// ErrNotEnoughCash ...
	ErrNotEnoughCash = errors.New("not enough cash")

	// ErrInvalidFieldForOperations ...
	ErrInvalidFieldForOperations = errors.New("invalid field for operation type")

	//ErrNotValidPhone
	ErrNotValidPhone = errors.New("invalid field for phone type")

	//ErrNotValidFirstName
	ErrNotValidFirstName = errors.New("invalid field for firstname type")

	//ErrNotValidLastName
	ErrNotValidLastName = errors.New("invalid field for lastname type")

	//ErrorNotValidPassword
	ErrorNotValidPassword = errors.New("invalid field for Password")

	//ErrorSignInCorrect
	ErrorSignInCorrect = errors.New("username or password is incorrect")
)

// errorHandler function for handling errors in product service
func ErrorHandler(err error, message string) error {
	if err == nil {
		return nil
	} else if strings.Contains("no rows in result set ", err.Error()) {
		return status.Error(codes.NotFound, ErrNotFound.Error())
	} else if strings.Contains("duplicate key value", err.Error()) {
		return status.Error(codes.AlreadyExists, message)
	} else if strings.Contains(err.Error(), "violates foreign key constraint") {
		return status.Error(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return status.Error(codes.Internal, InternalServerError)
	} else {
		return status.Error(codes.Unknown, err.Error())
	}

}

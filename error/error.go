package error

import (
	"fmt"
)

type AuthSignatureInvalid struct {}
func (AuthSignatureInvalid) Error() string {
	return "auth signature is invalid"
}

type AuthFailed struct {}
func (AuthFailed) Error() string {
	return "auth failed"
}

type TokenExpired struct {}
func (TokenExpired) Error() string {
	return "token has expired"
}

type TokenInvalid struct {}
func (TokenInvalid) Error() string {
	return "token is invalid"
}

type TokenSigningError struct {}
func (TokenSigningError) Error() string {
	return "token signing error"
}

type UserAlreadyExists struct {
	Name string `json:"name"`
}
func (e UserAlreadyExists) Error() string {
	return fmt.Sprintf("user with name %s already exists", e.Name)
}

type UserNotFound struct {
	Name string `json:"name"`
}
func (e UserNotFound) Error() string {
	return fmt.Sprintf("user with name %s is not found", e.Name)
}

type InvalidUserId struct {
	ID string `json:"id"`
}
func (e InvalidUserId) Error() string {
	return fmt.Sprintf("invalid user id: %s", e.ID)
}

type DbTransactionFailed struct {
	Err error
}
func (e DbTransactionFailed) Error() string {
	return fmt.Sprintf("db transaction error: %s", e.Err)
}

type InternalServerError struct {
	Err string `json:"error"`
}
func (e InternalServerError) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

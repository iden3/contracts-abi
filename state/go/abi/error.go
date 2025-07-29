package abi

import (
	"errors"
	"strings"
)

var (
	msgErrIdentityDoesNotExist            = "Identity does not exist"
	msgErrIdentityDoesNotExistCustomError = "IdentityDoesNotExist"
	msgErrRootDoesNotExist                = "Root does not exist"
	msgErrRootDoesNotExistCustomError     = "RootDoesNotExist"
	msgErrStateDoesNotExist               = "State does not exist"
	msgErrStateDoesNotExistCustomError    = "StateDoesNotExist"
)

var (
	ErrIdentityDoesNotExist = errors.New(msgErrIdentityDoesNotExist)
	ErrRootDoesNotExist     = errors.New(msgErrRootDoesNotExist)
	ErrStateDoesNotExist    = errors.New(msgErrStateDoesNotExist)
)

func isError(err error, errorMessage string) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), errorMessage)

}

func IsErrIdentityDoesNotExist(err error) bool {
	return isError(err, msgErrIdentityDoesNotExist) || isError(err, msgErrIdentityDoesNotExistCustomError)
}

func IsErrRootDoesNotExist(err error) bool {
	return isError(err, msgErrRootDoesNotExist) || isError(err, msgErrRootDoesNotExistCustomError)
}

func IsErrStateDoesNotExist(err error) bool {
	return isError(err, msgErrStateDoesNotExist) || isError(err, msgErrStateDoesNotExistCustomError)
}

package abi

import (
	"errors"
	"strings"
)

var (
	msgErrNodeNotFound                 = "Node not found"
	msgErrNodeNotFoundCustomError      = "NodeNotFound"
	msgErrInvalidStateNode             = "Invalid state node"
	msgErrInvalidStateNodeCustomError  = "InvalidStateNode"
	msgErrInvalidNodeType              = "Invalid node type"
	msgErrInvalidNodeTypeCustomError   = "InvalidNodeType"
	msgErrUnsupportedLength            = "Unsupported length"
	msgErrUnsupportedLengthCustomError = "UnsupportedLength"
	msgErrInvalidRootsLength           = "Invalid roots length"
)

var (
	ErrNodeNotFound      = errors.New(msgErrNodeNotFound)
	ErrInvalidStateNode  = errors.New(msgErrInvalidStateNode)
	ErrInvalidNodeType   = errors.New(msgErrInvalidNodeType)
	ErrUnsupportedLength = errors.New(msgErrUnsupportedLength)
)

func isError(err error, errorMessage string) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), errorMessage)

}

func IsErrNodeNotFound(err error) bool {
	return isError(err, msgErrNodeNotFound) || isError(err, msgErrNodeNotFoundCustomError)
}

func IsErrInvalidStateNode(err error) bool {
	return isError(err, msgErrInvalidStateNode) || isError(err, msgErrInvalidStateNodeCustomError) || isError(err, msgErrInvalidRootsLength)
}

func IsErrInvalidNodeType(err error) bool {
	return isError(err, msgErrInvalidNodeType) || isError(err, msgErrInvalidNodeTypeCustomError)
}

func IsErrUnsupportedLength(err error) bool {
	return isError(err, msgErrUnsupportedLength) || isError(err, msgErrUnsupportedLengthCustomError)
}

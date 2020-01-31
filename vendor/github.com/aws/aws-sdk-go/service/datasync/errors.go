// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalException for service response error code
	// "InternalException".
	//
	// This exception is thrown when an error occurs in the AWS DataSync service.
	ErrCodeInternalException = "InternalException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// This exception is thrown when the client submits a malformed request.
	ErrCodeInvalidRequestException = "InvalidRequestException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalException":       newErrorInternalException,
	"InvalidRequestException": newErrorInvalidRequestException,
}

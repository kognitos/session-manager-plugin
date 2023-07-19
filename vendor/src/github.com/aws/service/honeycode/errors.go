// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package honeycode

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action. Check that the
	// workbook is owned by you and your IAM policy allows access to the resource
	// in the request.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAutomationExecutionException for service response error code
	// "AutomationExecutionException".
	//
	// The automation execution did not end successfully.
	ErrCodeAutomationExecutionException = "AutomationExecutionException"

	// ErrCodeAutomationExecutionTimeoutException for service response error code
	// "AutomationExecutionTimeoutException".
	//
	// The automation execution timed out.
	ErrCodeAutomationExecutionTimeoutException = "AutomationExecutionTimeoutException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// There were unexpected errors from the server.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeRequestTimeoutException for service response error code
	// "RequestTimeoutException".
	//
	// The request timed out.
	ErrCodeRequestTimeoutException = "RequestTimeoutException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// A Workbook, Table, App, Screen or Screen Automation was not found with the
	// given ID.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The request caused service quota to be breached.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// Remote service is unreachable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Tps(transactions per second) rate reached.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Request is invalid. The message in the response contains details on why the
	// request is invalid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":               newErrorAccessDeniedException,
	"AutomationExecutionException":        newErrorAutomationExecutionException,
	"AutomationExecutionTimeoutException": newErrorAutomationExecutionTimeoutException,
	"InternalServerException":             newErrorInternalServerException,
	"RequestTimeoutException":             newErrorRequestTimeoutException,
	"ResourceNotFoundException":           newErrorResourceNotFoundException,
	"ServiceQuotaExceededException":       newErrorServiceQuotaExceededException,
	"ServiceUnavailableException":         newErrorServiceUnavailableException,
	"ThrottlingException":                 newErrorThrottlingException,
	"ValidationException":                 newErrorValidationException,
}
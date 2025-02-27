// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Another modification is being made. That modification must complete before you
// can make your change.
type ConcurrentModificationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ConcurrentModificationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The next token is not valid.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service role is not valid.
type InvalidServiceRoleException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidServiceRoleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidServiceRoleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidServiceRoleException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidServiceRoleException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidServiceRoleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource limit has been exceeded.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An AWS CodeStar project with the same ID already exists in this region for the
// AWS account. AWS CodeStar project IDs must be unique within a region for the AWS
// account.
type ProjectAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ProjectAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProjectAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProjectAlreadyExistsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ProjectAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ProjectAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Project configuration information is required but not specified.
type ProjectConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ProjectConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProjectConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProjectConfigurationException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ProjectConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ProjectConfigurationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The project creation request was valid, but a nonspecific exception or error
// occurred during project creation. The project could not be created in AWS
// CodeStar.
type ProjectCreationFailedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ProjectCreationFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProjectCreationFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProjectCreationFailedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ProjectCreationFailedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ProjectCreationFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified AWS CodeStar project was not found.
type ProjectNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ProjectNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProjectNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProjectNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ProjectNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ProjectNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The team member is already associated with a role in this project.
type TeamMemberAlreadyAssociatedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TeamMemberAlreadyAssociatedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TeamMemberAlreadyAssociatedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TeamMemberAlreadyAssociatedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "TeamMemberAlreadyAssociatedException"
	}
	return *e.ErrorCodeOverride
}
func (e *TeamMemberAlreadyAssociatedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified team member was not found.
type TeamMemberNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TeamMemberNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TeamMemberNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TeamMemberNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "TeamMemberNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *TeamMemberNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A user profile with that name already exists in this region for the AWS account.
// AWS CodeStar user profile names must be unique within a region for the AWS
// account.
type UserProfileAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UserProfileAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserProfileAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserProfileAlreadyExistsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "UserProfileAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *UserProfileAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user profile was not found.
type UserProfileNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UserProfileNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserProfileNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserProfileNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "UserProfileNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *UserProfileNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified input is either not valid, or it could not be validated.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

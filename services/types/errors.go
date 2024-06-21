package types

import "fmt"

//returned if the request if the malformed or contains an error
type BadequestExeption struct{
	Message *string
	ErrorCOdeOverride *string
	code *string
	Errorfault *string
	noSmithDocuments
}

type Errorfault struct{
	smith error
}

type noSmithDocuments error

func (e *BadequestExeption) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}

func (e *BadequestExeption) ErrorCode() string{
	if e ==nil || e.ErrorCOdeOverride == nil{
		return "BadRequestExeption"
	}
	return *e.ErrorCOdeOverride
}

func (e *BadequestExeption) ErrorMessage() string{
	if e.Message == nil{
		return ""
	}
	return *e.Message
}

func (e *BadequestExeption) ErrorFault() Errorfault{
	return Errorfault{
		smith: e.ErrorFault().smith,
	}
}

//returned if the session doesn't exist anumore because it timed out or expired 
type InvalidSessionException struct{
	Message *string
	ErrorCOdeOverride *string
	Code *string
	noSmithDocuments
}

func (e *InvalidSessionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSessionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSessionException) ErrorCode() string {
	if e == nil || e.ErrorCOdeOverride == nil {
		return "InvalidSessionException"
	}
	return *e.ErrorCOdeOverride
}

func (e *InvalidSessionException) ErrorFault() Errorfault{
	return Errorfault{
		smith: e.ErrorFault().smith,
	}
}

//returned if the session doesn't exist anumore because it timed out or expired 
type LimitExceededException struct{
	Message *string
	ErrorCOdeOverride *string
	Code *string
	noSmithDocuments
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
	if e == nil || e.ErrorCOdeOverride == nil {
		return "InvalidSessionException"
	}
	return *e.ErrorCOdeOverride
}

func (e *LimitExceededException) ErrorFault() Errorfault{
	return Errorfault{
		smith: e.ErrorFault().smith,
	}
}

// Returned when a transaction cannot be written to the journal due to a failure
// in the verification phase of optimistic concurrency control (OCC).
type OccConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithDocuments
}

func (e *OccConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OccConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OccConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OccConflictException"
	}
	return *e.ErrorCodeOverride
}

func (e *OccConflictException) ErrorFault() Errorfault{
	return Errorfault{
		smith: e.ErrorFault().smith,
	}
}

// Returned when the rate of requests exceeds the allowed throughput.
type RateExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithDocuments
}

func (e *RateExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RateExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RateExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RateExceededException"
	}
	return *e.ErrorCodeOverride
}

func (e *RateExceededException) ErrorFault() Errorfault{
	return Errorfault{
		smith: e.ErrorFault().smith,
	}
}
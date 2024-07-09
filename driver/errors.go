package driver

type txnError struct {
	transactionID string
	message       string
	err           error
	canRetry      bool
	abortSuccess  bool
	isISE         bool
}

// qldbDriverError is returned when an error caused by QLDBDriver has occurred.
type eqldbDriverError struct {
	errorMessage string
}

// Return the message denoting the cause of the error.
func (e *eqldbDriverError) Error() string {
	return e.errorMessage
}

func (e *txnError) unwrap() error {
	return e.err
}
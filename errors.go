package encryptequantumledgerdatabase

//EQLDB

// eqldbError is returned when an error caused by EQLDB has occurred.
type eqldbError struct {
	errormessage string
}

// Return the message denoting the cause of the error
func (e *eqldbError) Error() string {
	return e.errormessage
}

type TxnError struct {
	transactionID string
	message       string
	err           error
	canRetry      bool
	abortSuccess  bool
	isISE         bool
}

func (e *TxnError) Unwrap() error {
	return e.err
}

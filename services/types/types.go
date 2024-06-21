package types

type ABortTransactionRequest struct {
	noSmithDocuments
}

type TimingInformation struct {

	// The amount of time that QLDB spent on processing the command, measured in
	// milliseconds.
	ProcessingTimeMilliseconds int64

	noSmithyDocumentSerde
}

// Contains the detals of the transaction to abort.
type ABortTransactionResult struct {
	//contains server-side performace information
	TimingInformation *TimingInformation
	noSmithDocuments
}

// Contains the details of the transaction to commit.
type CommitTransctionRequest struct {
	CommitDigest  []byte
	TransactionID *string
	noSmithDocuments
}

// Contains the details of the committed transaction.
type CommitTransactionResult struct {
	// The commit digest of the committed transaction.
	CommitDigest []byte

	// Contains metrics about the number of I/O requests that were consumed.
	ConsumedIOs *IOUsage

	// Contains server-side performance information for the command.
	TimingInformation *TimingInformation

	// The transaction ID of the committed transaction.
	TransactionId *string
	noSmithDocuments
}

// Specifies a request to end the session.
type EndSessionRequest struct {
	noSmithDocuments
}

type EndSessionResult struct {

	// Contains server-side performance information for the command.
	TimingInformation *TimingInformation
	noSmithyDocumentSerde
}

// Specifies a request to execute a statement.
type ExecuteStatementRequest struct {

	// Specifies the statement of the request.
	//
	// This member is required.
	Statement *string

	// Specifies the transaction ID of the request.
	//
	// This member is required.
	TransactionId *string

	// Specifies the parameters for the parameterized statement in the request.
	Parameters []ValueHolder

	noSmithyDocumentSerde
}
type noSmithyDocumentSerde = noSmithDocuments

type NoSerde struct{}

func (n NoSerde) noSmithyDocumentSerde() {}

//type noSmithyDocumentSerde = noSmithDocuments.NoSerde

// Contains the details of the executed statement.
type ExecuteStatementResult struct {

	// Contains metrics about the number of I/O requests that were consumed.
	ConsumedIOs *IOUsage

	// Contains the details of the first fetched page.
	FirstPage *Page

	// Contains server-side performance information for the command.
	TimingInformation *TimingInformation

	noSmithyDocumentSerde
}

// Specifies the details of the page to be fetched.
type FetchPageRequest struct {

	// Specifies the next page token of the page to be fetched.
	//
	// This member is required.
	NextPageToken *string

	// Specifies the transaction ID of the page to be fetched.
	//
	// This member is required.
	TransactionId *string

	noSmithyDocumentSerde
}

// Contains the page that was fetched.
type FetchPageResult struct {

	// Contains metrics about the number of I/O requests that were consumed.
	ConsumedIOs *IOUsage

	// Contains details of the fetched page.
	Page *Page

	// Contains server-side performance information for the command.
	TimingInformation *TimingInformation

	noSmithyDocumentSerde
}

// Contains I/O usage metrics for a command that was invoked.
type IOUsage struct {

	// The number of read I/O requests that the command made.
	ReadIOs int64

	// The number of write I/O requests that the command made.
	WriteIOs int64

	noSmithyDocumentSerde
}

// Contains details of the fetched page.
type Page struct {

	// The token of the next page.
	NextPageToken *string

	// A structure that contains values in multiple encoding formats.
	Values []ValueHolder

	noSmithyDocumentSerde
}

// Specifies a request to start a new session.
type StartSessionRequest struct {

	// The name of the ledger to start a new session against.
	//
	// This member is required.
	LedgerName *string

	noSmithyDocumentSerde
}

// Specifies a request to start a transaction.
type StartTransactionRequest struct {
	noSmithyDocumentSerde
}

// Contains the details of the started transaction.
type StartTransactionResult struct {

	// Contains server-side performance information for the command.
	TimingInformation *TimingInformation

	// The transaction ID of the started transaction.
	TransactionId *string

	noSmithyDocumentSerde
}

// Contains the details of the started session.
type StartSessionResult struct {

	// Session token of the started session. This SessionToken is required for every
	// subsequent command that is issued during the current session.
	SessionToken *string

	// Contains server-side performance information for the command.
	TimingInformation *TimingInformation

	noSmithyDocumentSerde
}

// A structure that can contain a value in multiple encoding formats.
type ValueHolder struct {

	// An Amazon Ion binary value contained in a ValueHolder structure.
	IonBinary []byte

	// An Amazon Ion plaintext value contained in a ValueHolder structure.
	IonText *string

	noSmithyDocumentSerde
}

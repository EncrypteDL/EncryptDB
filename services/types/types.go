package types

type (
	noSerde interface {
		noSmithyDocumentSerde(error)
	}
	smithydocument error
)

type noSmithyDocumentSerde = smithydocument

// Contains the details of the transaction to abort.
type AbortTransactionrequest struct {
	noSmithyDocumentSerde
}

// Contains The detail of teh abord transaction.
type ABortTransaction struct {

	//Conatains server-side performance information for the command
	TimingInformation *TimingInformation
	noSmithyDocumentSerde
}

//type TimingInformation time.Time

// Contains the details of the transaction to commit.
type CommitTransactionRequest struct {
	CommitDigest  []byte
	TransactionID *string
	noSmithyDocumentSerde
}

// Contains the details of the \Commited transaction
type CommitTrasactionResult struct {
	//The commit digest of the commited transaction
	CommitDigest []byte

	//Contains metrics about the number of I/O request that were consumed
	ConsumedIOs *IOUsage

	// Contains server-side performance information for the command.
	TimingInformation *TimingInformation

	// The transaction ID of the committed transaction.
	TransactionId *string

	noSmithyDocumentSerde
}

// Specific a request to end the session.
type EndSessionREquest struct {
	noSmithyDocumentSerde
}

// Contains the details of the ended session.
type ENdSessionResult struct {
	//Contains server-side performance information for the command
	TimingInformation *TimingInformation
	noSmithyDocumentSerde
}

// Speciies a request to execute a statement
type ExecuteStatementRequest struct {
	Statement     *string
	TransactionID *string

	// Specifies the parameters for the parameterized statement in the request.
	Parameters []ValueHolder

	noSmithyDocumentSerde
}

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
	//SPecifie the nest page token of the page to be fetched
	NextPageToken *string

	//Specifies transaction ID of the page to be fetched
	Transactionid *string
	noSmithyDocumentSerde
}

// Contains the page that was fetched.
type FetchPageResult struct {
	ConsumedIOs *IOUsage
	// Contains details of the fetched page.
	Page *Page

	TimingInformation *TimingInformation

	noSmithyDocumentSerde
}

// COnatains I.O usage metrics for a command that was invoked.
type IOUsage struct {
	//The number of read I/Orequest that the command made
	ReadIOS int64

	//The num of write I/O request tahta the command line
	WriteIOs int64
	noSmithyDocumentSerde
}

// Conatins details of the page.
type Page struct {
	// The token of the next page.
	NextPageToken *string

	// A structure that contains values in multiple encoding formats.
	Values []ValueHolder

	noSmithyDocumentSerde
}

// Specifies a rrequest to strat a new session
type StartSessionRequest struct {
	// Session token of the started session. This SessionToken is required for every
	// subsequent command that is issued during the current session.
	SessionToken *string

	//COntains server-side performance information for the command .
	TimingInformation *TimingInformation
	noSmithyDocumentSerde
}

// Specifies a request to start a transaction.
type StartTransactionRequest struct {
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

// Contains the details of the stated transaction.
type StartTransactionResult struct {
	//contain server-side performance information for the command
	TimingInformation *TimingInformation

	//The Transaction ID of the started trnasaction
	TransactionID *string
	noSmithyDocumentSerde
}

// Contains server-side performance information for a command. Amazon QLDB
// captures timing information between the times when it receives the request and
// when it sends the corresponding response.
type TimingInformation struct {

	// The amount of time that QLDB spent on processing the command, measured in
	// milliseconds.
	ProcessingTimeMilliseconds int64

	noSmithyDocumentSerde
}

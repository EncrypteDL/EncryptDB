package interfac

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/qldbsession/types"
	"github.com/aws/smithy-go/middleware"
)

type SendCommandInput struct {

	// Command to abort the current transaction.
	AbortTransaction *types.AbortTransactionRequest

	// Command to commit the specified transaction.
	CommitTransaction *types.CommitTransactionRequest

	// Command to end the current session.
	EndSession *types.EndSessionRequest

	// Command to execute a statement in the specified transaction.
	ExecuteStatement *types.ExecuteStatementRequest

	// Command to fetch a page.
	FetchPage *types.FetchPageRequest

	// Specifies the session token for the current command. A session token is
	// constant throughout the life of the session.
	//
	// To obtain a session token, run the StartSession command. This SessionToken is
	// required for every subsequent command that is issued during the current session.
	SessionToken *string

	// Command to start a new session. A session token is obtained as part of the
	// response.
	StartSession *types.StartSessionRequest

	// Command to start a new transaction.
	StartTransaction *types.StartTransactionRequest

	//noSmithyDocumentSerde
}

type SendCommandOutput struct {

	// Contains the details of the aborted transaction.
	AbortTransaction *types.AbortTransactionResult

	// Contains the details of the committed transaction.
	CommitTransaction *types.CommitTransactionResult

	// Contains the details of the ended session.
	EndSession *types.EndSessionResult

	// Contains the details of the executed statement.
	ExecuteStatement *types.ExecuteStatementResult

	// Contains the details of the fetched page.
	FetchPage *types.FetchPageResult

	// Contains the details of the started session that includes a session token. This
	// SessionToken is required for every subsequent command that is issued during the
	// current session.
	StartSession *types.StartSessionResult

	// Contains the details of the started transaction.
	StartTransaction *types.StartTransactionResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	//noSmithyDocumentSerde
}

type ClientAPI interface {
	SendCommand(ctx context.Context, params *SendCommandInput, optFns ...func(Options)) (SendCommandOutput, error)
}

type Options struct {
	Retryer time.Duration
}

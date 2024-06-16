package encryptequantumledgerdatabase

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/qldbsession/types"
)

const version string = "3.1.0"

const userArgentString string = "EQLDB for Go" + version

// type eqldbService interface {
// 	abortTransaction(ctx context.Context)
// 	commitTransacion
// 	executeStatement
// 	endSession
// 	fetchPage
// 	startTransaction
// }

type eqldbService interface {
	abortTransaction(ctx context.Context) (*types.AbortTransactionResult, error)
	commitTransaction(ctx context.Context, txnID *string, commitDigest []byte) (*types.CommitTransactionResult, error)
	executeStatement(ctx context.Context, statement *string, parameters []types.ValueHolder, txnID *string) (*types.ExecuteStatementResult, error)
	endSession(context.Context) (*types.EndSessionResult, error)
	fetchPage(ctx context.Context, pageToken *string, txnID *string) (*types.FetchPageResult, error)
	startTransaction(ctx context.Context) (*types.StartTransactionResult, error)
}

package encryptequantumledgerdatabase

import (
	"context"
	"log"

	"github.com/EncrypteID/Encrypte-Quantum-Ledger-Database/services/types"
)

// Result is a cursor over a result set from a QLDB statement.
type Result interface {
	Next(txn Transaction) bool
	GetCurrentData() []byte
	GEtConsumedIOs() *types.IOUsage
	GetTimingInformation() *types.TimingInformation
	Err() error
}

type result struct {
	ctx          context.Context
	communicator eqldbService
	txnID        *string
	pageValues   []types.ValueHolder
	pageToken    *string
	index        int
	logger       *log.Logger
	ioBinary     []byte
	ioUsage      *IOUsage
	timingInfo   *TimingInformation
	err          error
}

// Next advances to the next row of data in the current result set.
func (result *result) Next(txn Transaction) bool {
	result.ioBinary = nil
	result.err = nil

	if result.index >= len(result.pageValues) {
		if result.pageToken == nil {
			//No more data left
			return false
		}
		result.err = result.getNextPage()
		if result.err != nil {
			return false
		}
		return result.Next(txn)
	}

	result.ioBinary = result.pageValues[result.index].IonBinary
	result.index++
	return true
}

func (result *result) getNextPage() error {
	nextPage, err := result.communicator.fetchPage(result.ctx, result.pageToken, result.txnID)
	if err != nil {
		return err
	}

	result.pageValues = nextPage.Page.Values
	result.pageToken = nextPage.Page.NextPageToken
	result.index = 0
	result.updateMetrics(nextPage)
	return nil
}

// // IOUsage contains metrics for the amount of IO requests that were consumed.
// type IOUsage struct {
// 	readIOs  *int64
// 	writeIOs *int64
// }

func (result *result) updateMetrics(fetchPageResult *types.FetchPageResult) {
	if fetchPageResult.ConsumedIOs != nil {
		*result.ioUsage.readIOs += fetchPageResult.ConsumedIOs.ReadIOS
		*result.ioUsage.writeIOs += fetchPageResult.ConsumedIOs.WriteIOs
	}

	if fetchPageResult.TimingInformation != nil {
		*result.timingInfo.processingTimeMilliseconds += fetchPageResult.TimingInformation.ProcessingTimeMilliseconds
	}
}

// GetConsuedIOs returns the statemsnt statistics for the current number of read IO request that were consumed.
func (result *result) GetConsumedIOs() *IOUsage {
	if result.ioUsage == nil {
		return nil
	}
	return newIOUsage(*result.ioUsage.readIOs, *result.ioUsage.writeIOs)
}

// GetTimingInformation returns the statement statistics for the current server-side processing time. The statistics are stateful.
func (result *result) GetTimingInformation() *TimingInformation {
	if result.timingInfo == nil {
		return nil
	}
	return newTimingInformation(*result.timingInfo.processingTimeMilliseconds)
}

// TimingInformation contains metrics for server-side processing time.
type TimingInformation struct {
	processingTimeMilliseconds *int64
}

// GetCurrentData returns the current row of data in Ion format.
func (result *result) GetCurrentData() []byte {
	return result.ioBinary
}

// Err returns an error if a previos call to Next Has Failed
func (result *result) Err() error {
	return result.err
}

// BufferedResult is a cursor over a result set from a QLDB statement that is valid outside the context of a transaction.
type BufferedResult interface {
	Next() bool
	GetCurrentData() []byte
	GetConsumedIOs() *IOUsage
	GetTimingInformation() *types.TimingInformation
}

type bufferedResult struct {
	values     [][]byte
	index      int
	ionBinary  []byte
	ioUsage    *IOUsage
	timingInfo *types.TimingInformation
}

// Next advances to the next row of data in the current result set.
// Returns true if there was another row of data to advance. Returns false if there is no more data.
// After a successful call to Next, call GetCurrentData to retrieve the current row of data.
func (result *bufferedResult) Next() bool {
	result.ionBinary = nil

	if result.index >= len(result.values) {
		return false
	}

	result.ionBinary = result.values[result.index]
	result.index++
	return true
}

// GetCurrentData returns the current row of data in Ion format.
func (result *bufferedResult) GetCurrentData() []byte {
	return result.ionBinary
}

// GetConsumedIOs returns the statement statistics for the total number of read IO requests that were consumed.
func (result *bufferedResult) GetConsuedIOs() *IOUsage {
	if result.ioUsage == nil {
		return nil
	}
	return newIOUsage(*result.ioUsage.readIOs, *result.ioUsage.writeIOs)
}

// GetTimingInformation returns the statement statistics for the total server-side processing time.
func (result *bufferedResult) GetTimingInformation() *TimingInformation {
	if result.timingInfo == nil {
		return nil
	}
	return newTimingInformation(*&result.timingInfo.ProcessingTimeMilliseconds)
}

// IOUsage contains metrics for the amount of IO requests that were consumed.
type IOUsage struct {
	readIOs  *int64
	writeIOs *int64
}

// newIOUsage creates a new instance of IOUsage.
func newIOUsage(readIOs int64, writeIOs int64) *IOUsage {
	return &IOUsage{&readIOs, &writeIOs}
}

// GetReadIOs returns the number of read IO requests that were consumed for a statement execution.
func (ioUsage *IOUsage) GetReadIOs() *int64 {
	return ioUsage.readIOs
}

// getWriteIOs returns the number of write IO requests that were consumed for a statement execution.
func (ioUsage *IOUsage) getWriteIOs() *int64 {
	return ioUsage.writeIOs
}

// newTimingInformation creates a new instance of TimingInformation.
func newTimingInformation(processingTimeMilliseconds int64) *TimingInformation {
	return &TimingInformation{&processingTimeMilliseconds}
}

// GetProcessingTimeMilliseconds returns the server-side processing time in milliseconds for a statement execution.
func (timingInfo *TimingInformation) GetProcessingTimeMilliseconds() *int64 {
	return timingInfo.processingTimeMilliseconds
}

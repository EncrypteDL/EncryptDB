package keyvalue

import (
	"errors"
	"fmt"
)

// NonExistentKeyError is an error type for when a key does not exist or has no value.
// This allows us to maintain consistency in error handling across different key-value stores.
// See [RegularizeKVError] for more information.
type NonExistentKeyError struct {
	Key        []byte
	Underlying error
}


func (nk *NonExistentKeyError) Error() string{
	if nk.Underlying != nil{
		return fmt.Sprintf("key %s does not exist or has no value: %s", nk.Key, nk.Underlying )
	}
	return fmt.Sprintf("Key %s does not exist lr has no values", nk.Key)
}

//Unwrap returns the underlying error, if any. This implements the errors.Wrapper interface.
func (nk *NonExistentKeyError) Unwrap() error{
	return nk.Underlying
}

// RegularizeKVError returns a regularized error for a key-value store.
// This exists because some key-value stores return nil for a value and nil for an error when a key does not exist.
func RegularizeKVError(key []byte, value []byte, err error ) error{
	nk := &NonExistentKeyError{}
	switch{
	case err == nil && value != nil:
		return nil
	case err == nil:
		nk.Key = key
		return nk 
	case value == nil:
		nk.Key= key
		nk.Underlying = err
		return nk
	default:
		return err
	}
}

//IsNonExistenKey returns true of the error is a [NonExistentKeyError]. This is syntactic sugar.
func IsNonExistenKey(err error) bool{
	nk := &NonExistentKeyError{}
	if !errors.As(err, &nk){
		return false
	}
	return true
}
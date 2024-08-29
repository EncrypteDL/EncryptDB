package encryptdb

import keyvalue "github.com/EncrypteDL/EncryptDB/key_value"

// Searcher must be able to search through our datastore(s) with strings.
type Searcher interface {
	// PrefixScan must retrieve all keys in the datastore and stream them to the given channel.
	PrefixScan(prefix string) (<-chan keyvalue.KeyValue, chan error)
	// Search must be able to search through the value contents of our database and stream the results to the given channel.
	Search(query string) (<-chan keyvalue.KeyValue, chan error)
	// ValueExists searches for an exact match of the given value and returns the key that contains it.
	ValueExists(value []byte) (key []byte, ok bool)
}

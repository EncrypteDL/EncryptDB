package driver

import (
	"crypto/sha256"
	"hash"

	"github.com/amzn/ion-go/ion"
)

// Type indicates a standard hashing algorithm
type Type int

type hashDefinition struct {
	width    int
	name     string
	alias    string
	newFunc  func() hash.Hash
	hashType Type
}

type eqldbHash struct {
	hash []byte
}

var (
	type2hash  = map[Type]*hashDefinition{}
	name2hash  = map[string]*hashDefinition{}
	alias2hash = map[string]*hashDefinition{}
	supported  = []Type{}
)

const hashsize = 32

// RegisterHash adds a new Hash to the list and returns it Type
func RegisterHash(name, alias string, width int, newFunc func() hash.Hash) Type {
	hashType := Type(1 << len(supported))
	supported = append(supported, hashType)

	definition := &hashDefinition{
		name:     name,
		alias:    alias,
		width:    width,
		newFunc:  newFunc,
		hashType: hashType,
	}

	type2hash[hashType] = definition
	name2hash[name] = definition
	alias2hash[alias] = definition

	return hashType
}

// CustomHashReader defines a custom hash reader
type CustomHashReader struct {
	reader ion.Reader
	hasher hash.Hash
}

// NewCustomHashReader creates a new custom hash reader
func NewCustomHashReader(reader ion.Reader) *CustomHashReader {
	return &CustomHashReader{
		reader: reader,
		hasher: sha256.New(),
	}
}

// Next processes the next value in the reader
func (chr *CustomHashReader) Next() bool {
	if chr.reader.Next() {
		value, err := ion.MarshalText(chr.reader.Value())
		if err != nil {
			return false
		}
		chr.hasher.Write(value)
		return true
	}
	return false
}

// Sum returns the final hash value
func (chr *CustomHashReader) Sum(b []byte) ([]byte, error) {
	return chr.hasher.Sum(b), nil
}

func toEldbhash(value interface{}) (*eqldbHash, error) {
	ionValue, err := ion.MarshalBinary(value)
	if err != nil {
		return nil, err
	}
	ionReader := ion.NewReaderBytes(ionValue)
	hashReader := NewCustomHashReader(ionReader)
	for hashReader.Next() {
		// Read over value
	}
	hash, err := hashReader.Sum(nil)
	if err != nil {
		return nil, err
	}
	return &eqldbHash{hash}, nil
}

func joinHashesPairwise(h1 []byte, h2 []byte) ([]byte, error) {
	if len(h1) == 0 {
		return h2, nil
	}
	if len(h2) == 0 {
		return h1, nil
	}

	compare, err := hashComparator(h1, h2)
	if err != nil {
		return nil, err
	}

	var concatenated []byte
	if compare < 0 {
		concatenated = append(h1, h2...)
	} else {
		concatenated = append(h2, h1...)
	}
	return concatenated, nil
}

func hashComparator(h1 []byte, h2 []byte) (int16, error) {
	if len(h1) != hashsize || len(h2) != hashsize {
		return 0, &eqldbDriverError{"invalid hash"}
	}
	for i := range h1 {
		// Reverse index for little endianness
		index := hashsize - 1 - i

		// Handle byte being unsigned and overflow
		h1Int := int16(h1[index])
		h2Int := int16(h2[index])
		if h1Int > 127 {
			h1Int = 0 - (256 - h1Int)
		}
		if h2Int > 127 {
			h2Int = 0 - (256 - h2Int)
		}

		difference := h1Int - h2Int
		if difference != 0 {
			return difference, nil
		}
	}
	return 0, nil
}

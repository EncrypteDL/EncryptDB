package hashing

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// HashSize of array used to store hashes.  See Hash std pkg.
const HashSize = 32

// MaxHashStringSize is the maximum length of a Hash hash string.
const MaxHashStringSize = HashSize * 2 //64

// ErrHashStrSize describes an error that indicates the caller specified a hash
// string that has too many characters.
var ErrHashStrSize = fmt.Errorf("max hash string length is %v bytes", MaxHashStringSize)

/*
-   Common Uses of Hashing
Data Integrity: Verifying that data has not been altered.
Cryptography: Ensuring secure data transmission and storage.
Hash Tables: Efficiently indexing and retrieving data in databases.
Digital Signatures: Authenticating the sender of a message.
Password Storage: Storing passwords in a hashed format to secure them.
*/

type Hash [HashSize]byte

// String returns the Hash as the hexadecimal string of the byte-reversed
// hash.
func (hash Hash) String() string {
	for i := 0; i < HashSize/2; i++ {
		hash[i], hash[HashSize-1-i] = hash[HashSize-1-i], hash[i]
	}
	return hex.EncodeToString(hash[:])
}

// CloneBytes sets the bytes which represent the hash, AN error is returned if
// the number of bytes passed in is not HashSize.
func (hash *Hash) CloneBytes() []byte {
	newHash := make([]byte, HashSize)
	copy(newHash, hash[:])

	return newHash
}

// SetBytes sets the bytes which represent the hash.  An error is returned if
// the number of bytes passed in is not HashSize.
func (hash *Hash) SetBytes(newHash []byte) error {
	nhlen := len(newHash)
	if nhlen != HashSize {
		return fmt.Errorf("invalid hash length of %v, want %v", nhlen,
			HashSize)
	}
	copy(hash[:], newHash)

	return nil
}

// IsEqual returns true if target is the same as hash.
func (hash *Hash) IsEqual(target *Hash) bool {
	if hash == nil && target == nil {
		return true
	}
	if hash == nil || target == nil {
		return false
	}
	return *hash == *target
}

// MArshal JSON serialises the ahsh as a JSON appropriate string value.
func (hash Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(hash.String())
}

// UnmarshalJSON parses the hash with JSON appropriate string value.
func (hash *Hash) UnmarshalJSON(input []byte) error {
	// If the first byte indicates an array, the hash could have been marshalled
	// using the legacy method and e.g. persisted.
	if len(input) > 0 && input[0] == '[' {
		return decodeLegacy(hash, input)
	}

	var sh string
	err := json.Unmarshal(input, &sh)
	if err != nil {
		return err
	}
	newHash, err := NewHashFromStr(sh)
	if err != nil {
		return err
	}

	return hash.SetBytes(newHash[:])
}

// NewHash returns a new Hash from a byte slice.  An error is returned if
// the number of bytes passed in is not HashSize.
func NewHash(newHash []byte) (*Hash, error) {
	var sh Hash
	err := sh.SetBytes(newHash)
	if err != nil {
		return nil, err
	}
	return &sh, err
}

// NewHashFromStr creates a Hash from a hash string.  The string should be
// the hexadecimal string of a byte-reversed hash, but any missing characters
// result in zero padding at the end of the Hash.
func NewHashFromStr(hash string) (*Hash, error) {
	ret := new(Hash)
	err := Decode(ret, hash)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Decode decodes the byte-reversed hexadecimal string encoding of a Hash to a
// destination.
func Decode(dst *Hash, src string) error {
	// Return error if hash string is too long.
	if len(src) > MaxHashStringSize {
		return ErrHashStrSize
	}

	// Hex decoder expects the hash to be a multiple of two.  When not, pad
	// with a leading zero.
	var srcBytes []byte
	if len(src)%2 == 0 {
		srcBytes = []byte(src)
	} else {
		srcBytes = make([]byte, 1+len(src))
		srcBytes[0] = '0'
		copy(srcBytes[1:], src)
	}

	// Hex decode the source bytes to a temporary destination.
	var reversedHash Hash
	_, err := hex.Decode(reversedHash[HashSize-hex.DecodedLen(len(srcBytes)):], srcBytes)
	if err != nil {
		return err
	}

	// Reverse copy from the temporary hash to destination.  Because the
	// temporary was zeroed, the written result will be correctly padded.
	for i, b := range reversedHash[:HashSize/2] {
		dst[i], dst[HashSize-1-i] = reversedHash[HashSize-1-i], b
	}

	return nil
}

// decodeLegacy decodes an Hash that has been encoded with the legacy method
// (i.e. represented as a bytes array) to a destination.
func decodeLegacy(dst *Hash, src []byte) error {
	var hashBytes []byte
	err := json.Unmarshal(src, &hashBytes)
	if err != nil {
		return err
	}
	if len(hashBytes) != HashSize {
		return ErrHashStrSize
	}
	return dst.SetBytes(hashBytes)
}

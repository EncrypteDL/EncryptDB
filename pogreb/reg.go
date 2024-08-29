package pogreb

import (
	encryptdb "github.com/EncrypteDL/EncryptDB"
	"github.com/EncrypteDL/EncryptDB/registry"
)

func init() {
	registry.RegisterKeeper("pogreb", func(path string, opts ...any) (encryptdb.Keeper, error) {
		if len(opts) > 1 {
			return nil, ErrInvalidOptions
		}
		if len(opts) == 1 {
			casted, castErr := castOptions(opts...)
			if castErr != nil {
				return nil, castErr
			}
			defOptMu.Lock()
			defaultPogrebOptions = casted
			defOptMu.Unlock()
		}
		db := OpenDB(path)
		err := db.init()
		return db, err
	})
}

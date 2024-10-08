package keeper

import (
	"fmt"
	"os"
	"path/filepath"

	encryptdb "github.com/EncrypteDL/EncryptDB"
	"github.com/EncrypteDL/EncryptDB/metadata"
	"github.com/EncrypteDL/EncryptDB/registry"
)

func OpenKeeper(path string, opts ...any) (encryptdb.Keeper, error) {
	stat, statErr := os.Stat(path)
	if statErr != nil {
		return nil, statErr
	}
	var metaDat []byte
	var readErr error
	if stat.IsDir() {
		if stat, statErr = os.Stat(filepath.Join(path, "meta.json")); statErr != nil {
			return nil, fmt.Errorf("meta.json not found in target directory: %w", os.ErrNotExist)
		}
		metaDat, readErr = os.ReadFile(filepath.Join(path, "meta.json"))
	} else {
		metaDat, readErr = os.ReadFile(path)
	}

	if readErr != nil {
		return nil, fmt.Errorf("error reading meta.json: %w", readErr)
	}
	if len(metaDat) == 0 {
		return nil, fmt.Errorf("meta.json is empty")
	}

	meta, err := metadata.LoadMeta(metaDat)
	if err != nil {
		return nil, fmt.Errorf("error parsing meta.json: %w", err)
	}
	var keeperCreator encryptdb.KeeperCreator
	if keeperCreator = registry.GetKeeper(meta.KeeperType); keeperCreator == nil {
		return nil, fmt.Errorf("keeper type %s not found in registry", meta.KeeperType)
	}

	var (
		keeper encryptdb.Keeper
	)

	if len(opts) > 0 {
		keeper, err = keeperCreator(path, opts...)
	} else {
		keeper, err = keeperCreator(path)
	}

	if err != nil {
		return nil, fmt.Errorf("error substantiating keeper: %w", err)
	}
	if _, err = keeper.Discover(); err != nil {
		return nil, fmt.Errorf("error opening existing stores: %w", err)
	}
	return keeper, nil
}

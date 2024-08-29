package pogreb

import (
	"testing"

	encryptdb "github.com/EncrypteDL/EncryptDB"
)

func Test_Interfaces(t *testing.T) {
	v := OpenDB(t.TempDir())
	var keeper interface{} = v
	if _, ok := keeper.(encryptdb.Keeper); !ok {
		t.Error("Keeper interface not implemented")
	} else {
		t.Log("Keeper interface implemented")
	}
	vs := v.WithNew("test")
	var searcher interface{} = vs
	if _, ok := searcher.(encryptdb.Searcher); !ok {
		t.Error("Searcher interface not implemented")
	} else {
		t.Log("Searcher interface implemented")
	}
	var filer interface{} = vs
	if _, ok := filer.(encryptdb.Filer); !ok {
		t.Error("Filer interface not implemented")
	} else {
		t.Log("Filer interface implemented")
	}
	var store *Store
	if !encryptdb.IsStore(store) {
		t.Error("Store interface not implemented")
	} else {
		t.Log("Store interface implemented")
	}
}

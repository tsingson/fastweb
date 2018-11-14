package txn

import (
	"errors"
	"github.com/dgraph-io/badger"
	"github.com/tsingson/fastweb/utils"
	"time"
)

const (
	TTL = 24 * time.Hour
)

type (
	Txn struct {
		Db *badger.DB
	}
)

func (txns *Txn) Close() {
	txns.Db.Close()
}
func NewTxn(path string) (*Txn, error) {
	if len(path) == 0 {
		return nil, errors.New("path do not exists")
	}
	var txn *Txn
	txn = new(Txn)

	opts := badger.DefaultOptions
	//	path, _ := fasthttputils.GetCurrentExecDir()
	opts.Dir = path + "/data"
	opts.ValueDir = path + "/data"
	db, err := badger.Open(opts)
	if err != nil {
		return txn, err
	}
	txn.Db = db
	return txn, nil
}

func (tsns *Txn) Set(key, value []byte) error {
	//key := uuid.NewV4().Bytes()
	//value := []byte("43")
	err := tsns.Db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		//	err := txn.SetWithTTL(key, value, TTL)
		return err
	})
	return err
}
func (tsns *Txn) SetWithTTL(key, value []byte, ttl time.Duration) error {
	//key := uuid.NewV4().Bytes()
	//value := []byte("43")
	err := tsns.Db.Update(func(txn *badger.Txn) error {
		//	err := txn.Set(key, value)
		err := txn.SetWithTTL(key, value, ttl)
		return err
	})
	return err
}

func (txns *Txn) GetStr(key []byte) (string, error) {
	var result []byte
	result, err := txns.Get(key)
	if err != nil {
		return "", err
	}

	return utils.BytesToStringUnsafe(result), nil
}

func (txns *Txn) Get(key []byte) ([]byte, error) {
	var result []byte
	err := txns.Db.View(func(txn *badger.Txn) error {

		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		val, err1 := item.Value()
		if err1 != nil {
			return err
		}
		result = val
		return err
	})
	return result, err
}
func (tsns *Txn) Delete(key []byte) error {
	// delete
	err := tsns.Db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
	return err
}

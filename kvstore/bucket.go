package kvstore

import (
	"sync"

	"github.com/dgraph-io/badger"
	"github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/tsingson/fastx/utils"
	"time"
)

type (
	Store struct {
		db  *badger.DB
		log zerolog.Logger
	}
	Bucket struct {
		db             *badger.DB
		ValidForPrefix []byte
		Name           string
		prefix         string
		log            zerolog.Logger
	}
)

var instance *Store
var once sync.Once

func Connect(db *badger.DB, log zerolog.Logger) *Store {
	once.Do(func() {
		instance = &Store{}
		instance.db = db
		sublogger := log.With().
			Str("module", "buket").
			Logger()
		sublogger.Level(zerolog.ErrorLevel)
		instance.log = sublogger
	})
	return instance
}

func (s *Store) Bucket(name string) *Bucket {
	bucket := new(Bucket)
	bucket.Name = name

	bucket.prefix = utils.StrBuilder(name, "-")
	bucket.ValidForPrefix = utils.StringToBytesUnsafe(bucket.prefix)
	bucket.db = s.db
	bucket.log = s.log
	return bucket
}

func (b *Bucket) Save(key string, val []byte) (err error) {

	txn := b.db.NewTransaction(true)
	defer txn.Discard()
	//
	id := utils.StrBuilder(b.prefix, key)
	keyByte := utils.StringToBytesUnsafe(id)
	b.log.Info().Str("key", id).Msg("bucket save")
	err2 := txn.Set(keyByte, val)

	if err2 != nil {
		//	b.log.Info().Err(err2).Msg("badger存入出错")
		return err2
	}

	// Commit the transaction and check for error.
	if err3 := txn.Commit(nil); err3 != nil {
		//	b.log.Info().Err(err3).Msg("badger事务提交出错")
		return err3
	}

	return nil
}

func (b *Bucket) SaveStruct(key string, val interface{}) (err error) {

	//	b.log.Info().Str("key", key).Msg("bucket save")
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	id := utils.StrBuilder(b.prefix, key)
	//	b.log.Info().Str("key", id).Msg("bucket save")
	keyByte := []byte(id)

	vodByte, err1 := jsoniter.Marshal(val)
	if err1 != nil {
		//	b.log.Info().Err(err1).Msg("json序列化出错")
		return err1
	}
	err2 := txn.Set(keyByte, vodByte)

	if err2 != nil {
		//	b.log.Info().Err(err2).Msg("badger存入出错")
		return err2
	}

	// Commit the transaction and check for error.
	if err3 := txn.Commit(nil); err3 != nil {
		//	b.log.Info().Err(err3).Msg("badger事务提交出错")
		return err3
	}
	b.log.Info().Str("key", id).Msg("bucket 存入成功")
	return nil
}
func (b *Bucket) SaveStructTTL(key string, val interface{}, ttl time.Duration) (err error) {

	b.log.Info().Str("key", key).Msg("bucket save")
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	id := utils.StrBuilder(b.prefix, key)
	keyByte := utils.StringToBytesUnsafe(id)

	vodByte, err1 := jsoniter.Marshal(val)
	if err1 != nil {
		b.log.Info().Err(err1).Msg("json序列化出错")
		return err1
	}
	err2 := txn.SetWithTTL(keyByte, vodByte, ttl)

	if err2 != nil {
		b.log.Info().Err(err2).Msg("badger存入出错")
		return err2
	}

	// Commit the transaction and check for error.
	if err3 := txn.Commit(nil); err3 != nil {
		b.log.Info().Err(err3).Msg("badger事务提交出错")
		return err3
	}

	return nil
}

func (b *Bucket) Fetch(key string) (val []byte, err error) {
	//	b.log.Info().Str("key", key).Msg("bucket fetch")

	id := utils.StrBuilder(b.prefix, key)
	//	b.log.Info().Str("key", id).Msg("bucket 获取的内容 Key")
	keyByte := []byte(id)
	//	var out []byte

	err0 := b.db.View(func(txn *badger.Txn) error {
		item, err1 := txn.Get(keyByte)
		if err1 != nil {
			//		b.log.Info().Err(err1).Msg("badger 获取的内容不存在")
			return err1
		}
		outByte, err2 := item.ValueCopy(nil)
		if err2 != nil {
			//		b.log.Info().Err(err2).Msg("badger ValueCopy 出错")
			return err2
		}
		val = outByte
		return nil
	})
	if err0 != nil {
		//	b.log.Info().Err(err0).Msg("badger 事务处理出错")
		return nil, err0
	}
	b.log.Info().Str("key", id).Msg("bucket 获取成功")
	return val, nil
}

func (b *Bucket) Delete(key string) (err error) {
	b.log.Info().Str("key", key).Msg("bucket delete")

	txn := b.db.NewTransaction(true)
	defer txn.Discard()
	//
	id := utils.StrBuilder(b.prefix, key)
	b.log.Info().Str("key", key).Msg("bucket delete")
	keyByte := []byte(id)
	//
	err2 := txn.Delete(keyByte)
	if err2 != nil {
		b.log.Info().Err(err2).Msg("badger 删除出错")
		txn.Discard()
		return err2
	}

	// Commit the transaction and check for error.
	if err3 := txn.Commit(nil); err3 != nil {
		b.log.Info().Err(err3).Msg("badger事务提交出错")
		txn.Discard()
		return err3
	}

	return nil
}

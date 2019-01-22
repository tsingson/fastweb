package txn

import (
	"github.com/sanity-io/litter"
	"github.com/tsingson/fastx/utils"
	"github.com/tsingson/uuid"
)

func txn_test() {
	path, _ := utils.GetCurrentExecDir()
	tsn, err := NewTxn(path)
	if err != nil {

	}
	defer tsn.Close()

	key := []byte("key")
	// set
	v, _ := uuid.NewV4()
	vvv := v.Bytes()
	err = tsn.Set(key, vvv)
	if err != nil {
		panic(err)
	}
	// get

	value, err1 := tsn.GetStr(key)
	if err1 != nil {
		panic(err)
	}
	litter.Dump(value)
	err = tsn.Delete(key)

	if err != nil {
		panic(err)
	}
	return
}

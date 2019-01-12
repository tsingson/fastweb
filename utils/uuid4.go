package utils

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	//"github.com/tsingson/btcutil/base58"
)

func Uuid4() string {
	uu, err1 := uuid.NewV4()
	if err1 != nil {

	}
	return hex.EncodeToString(uu.Bytes())
}

func Uuid4Base58() string {
	return base58.Encode(StringToBytesUnsafe(Uuid4()))
}

// design and code by tsingson

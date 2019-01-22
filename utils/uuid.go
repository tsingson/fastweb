package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/tsingson/btcutil/base58"
	"github.com/tsingson/fastx/guid"
	"github.com/tsingson/uuid"
)

type Guid struct {
	Uuid uuid.UUID
}

func NewGuid() string {
	//	enc := new(guid.Base58)
	uu, _ := uuid.NewV4()
	return hex.EncodeToString(uu.Bytes())
}

/**
func Uuid4() uuid.UUID {
	return uuid.NewV4()
}
*/
func MgoUuid() string {
	return guid.New128().String()
}
func (u Guid) String() string {
	return hex.EncodeToString(u.Uuid.Bytes())
}

func Bytes2String(u []byte) string {
	return hex.EncodeToString(u)
}
func base58Test() {

	u, _ := uuid.NewV4()

	uu := hex.EncodeToString(u.Bytes())

	ip := "192.168.1.1"

	fmt.Println("****************")
	res := base58.CheckEncode([]byte(ip), 20)
	fmt.Println("base58check", res)
	orgin, _, _ := base58.CheckDecode(res)
	fmt.Println("base58check decode ", string(orgin))

	fmt.Println("****************")
	uures := base58.Encode([]byte(uu))
	fmt.Println("base58", uures)
	uuorgin := base58.Decode(uures)
	fmt.Println("base58 decode ", string(uuorgin))

}

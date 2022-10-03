package common

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) *UID {
	return &UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

func (u UID) String() string {
	val := uint64(u.localID)<<28 | uint64(u.objectType)<<18 | uint64(u.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (u UID) GetLocationID() uint32 {
	return u.localID
}

func (u UID) GetObjectType() int {
	return u.objectType
}

func (u UID) GetShardID() uint32 {
	return u.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}

	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}

	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (u UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", u.String())), nil
}

func (u *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}

	u.localID = decodeUID.localID
	u.objectType = decodeUID.objectType
	u.shardID = decodeUID.shardID

	return nil
}

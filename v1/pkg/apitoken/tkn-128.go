package apitoken

import (
	"github.com/google/uuid"
	"strconv"
)

func NewToken128() Token128 {
	return Token128(uuid.New())
}

type Token128 [16]byte

func (t *Token128) FromString(tmp string) error {
	for i := range t {
		pos := i * 2
		if u, e := strconv.ParseUint(tmp[pos : pos+2], 16, 8); e != nil {
			return e
		} else {
			t[i] = byte(u)
		}
	}

	return nil
}

func (t Token128) String() string {
	return formatHexString(t[:])
}

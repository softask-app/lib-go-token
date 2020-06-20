package apitoken

import (
	"github.com/google/uuid"
	"strconv"
)

type Token256 [32]byte

func NewToken256() (out Token256) {
	front := uuid.New()
	back  := uuid.New()
	copy(out[0:], front[:])
	copy(out[16:], back[:])

	return
}

func (t Token256) String() string {
	return formatHexString(t[:])
}

func (t *Token256) FromString(tmp string) error {
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

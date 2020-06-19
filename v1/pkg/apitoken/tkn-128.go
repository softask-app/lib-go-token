package apitoken

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/google/uuid"
)

func NewToken128() Token128 {
	return Token128(uuid.New())
}

type Token128 [16]byte

func (t Token128) String() string {
	return formatHexString(t[:])
}

func (t Token128) MarshalJSONArray(enc *gojay.Encoder) {
	enc.String(t.String())
}

func (t Token128) IsNil() bool {
	return false
}

func (t Token128) UnmarshalJSONArray(d *gojay.Decoder) error {
	var tmp string
	if err := d.String(&tmp); err != nil {
		return err
	}

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

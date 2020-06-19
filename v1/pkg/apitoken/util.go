package apitoken

import (
	"strings"

	"github.com/vulpine-io/char/v1"
)

var hex = [...]byte{
	char.Zero,
	char.One,
	char.Two,
	char.Three,
	char.Four,
	char.Five,
	char.Six,
	char.Seven,
	char.Eight,
	char.Nine,
	char.LowerA,
	char.LowerB,
	char.LowerC,
	char.LowerD,
	char.LowerE,
	char.LowerF,
}

func formatHexString(bytes []byte) string {
	out := new(strings.Builder)
	out.Grow(len(bytes) * 2)
	out.Reset()

	for _, b := range bytes {
		writeHexByte(out, b)
	}

	return out.String()
}

func writeHexByte(w *strings.Builder, b byte) {
	w.WriteByte(hex[(b>>4)&0xF])
	w.WriteByte(hex[b&0xF])
}

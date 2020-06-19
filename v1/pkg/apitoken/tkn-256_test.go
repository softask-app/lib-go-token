package apitoken_test

import (
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewToken256(t *testing.T) {
	Convey("NewToken256", t, func() {
		a := apitoken.NewToken256()
		b := apitoken.NewToken256()

		So(a, ShouldNotEqual, b)
		So(a, ShouldNotResemble, b)
	})
}

func TestToken256_IsNil(t *testing.T) {
	Convey("Token256.IsNil", t, func() {
		So(apitoken.NewToken256().IsNil(), ShouldBeFalse)
	})
}

func TestToken256_String(t *testing.T) {
	Convey("Token256.String", t, func() {
		foo := apitoken.Token256{
			1,   5,  10,  50, 100, 255, 3,  13, 23, 56, 18,  99,  69,  8,    15, 9,
			200, 55, 188, 80, 90,  204, 78, 66, 86, 99, 200, 178, 145, 154,  19, 36,
		}

		So(foo.String(), ShouldEqual, "01050a3264ff030d1738126345080f09"+
			"c837bc505acc4e425663c8b2919a1324")
	})
}

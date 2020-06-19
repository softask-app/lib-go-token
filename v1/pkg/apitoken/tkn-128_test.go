package apitoken_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

func TestNewToken128(t *testing.T) {
	Convey("NewToken128", t, func() {
		a := apitoken.NewToken128()
		b := apitoken.NewToken128()

		So(a, ShouldNotEqual, b)
		So(a, ShouldNotResemble, b)
	})
}

func TestToken128_IsNil(t *testing.T) {
	Convey("Token128.IsNil", t, func() {
		So(apitoken.NewToken128().IsNil(), ShouldBeFalse)
	})
}

func TestToken128_String(t *testing.T) {
	Convey("Token128.String", t, func() {
		foo := apitoken.Token128{1, 5, 10, 50, 100, 255, 3, 13, 23, 56, 18, 99, 69, 8, 15, 9}

		So(foo.String(), ShouldEqual, "01050a3264ff030d1738126345080f09")
	})
}

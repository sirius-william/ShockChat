package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestUsing(t *testing.T) {
	Convey("Test Using str:\"hello world\"", t, func() {
		a, err1 := TestRSAEncrypt([]byte("hello world"))
		So(len(a), ShouldNotEqual, 0)
		b := TestRSADecrypt(a)
		So(err1, ShouldBeNil)
		So(b, ShouldEqual, "hello world")
	})
}

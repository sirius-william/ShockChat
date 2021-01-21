package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestKeyFileCreateSuccessfull(t *testing.T) {
	Convey("Test Key File Create Successfully.", t, func() {
		So(TestRSAGenerationOfKey(), ShouldBeTrue)
	})
}

func TestUsing(t *testing.T) {
	Convey("Test Using str:\"hello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello world\"", t, func() {
		a, err1 := TestRSAEncrypt([]byte("hello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello world"))
		So(len(a), ShouldNotEqual, 0)
		b := TestRSADecrypt(a)
		So(err1, ShouldBeNil)
		So(b, ShouldEqual, "hello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello world")
	})
}

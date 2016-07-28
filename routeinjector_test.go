package routeinjector

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSomething(t *testing.T) {
	Convey("The substraction should be zero", t, func() {
		So(1-1, ShouldEqual, 0)
	})
}

func TestSomethingElse(t *testing.T) {
	t.Log("Stupid test")
}

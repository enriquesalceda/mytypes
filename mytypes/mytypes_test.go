package mytypes_test

import (
	"mytypes/mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	var number mytypes.MyInt = 5

	var want mytypes.MyInt = 10
	got := number.Twice()

	if want != got {
		t.Errorf("Want: %d, Got: %d", want, got)
	}
}

func TestJoin(t *testing.T) {
	input := mytypes.MultiString{
		"a",
		"b",
		"c",
	}

	want := "abc"
	got := input.Join()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestHello(t *testing.T) {
	want := "Hello, Gophers!"
	var mb mytypes.MyBuilder
	got := mb.Hello()

	if want != got {
		t.Errorf("Want: %s, Got: %s", want, got)
	}
}

func TestToUp(t *testing.T) {
	var input mytypes.OtherBuilder
	input.Contents.WriteString("gopher")

	got := input.ToUp()
	want := "GOPHER"

	if want != got {
		t.Errorf("Want: %s, Got: %s", want, got)
	}
}

func TestSum(t *testing.T) {
	var input mytypes.IntBuilder
	input.Contents = []int{1, 2, 3}

	want := 6
	got := input.Sum()

	if want != got {
		t.Errorf("Want: %d, Got: %d", want, got)
	}
}

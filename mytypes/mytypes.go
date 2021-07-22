package mytypes

import "strings"

// EXTENDING INTEGER TYPE

// This is not possible because this is not part of the set of methods
// for int, we are trying to extende the definition, Go does not allow
// this, because this is not our type

// func (i int) Twice() {
// 	return i * 2
// }

// We can create our own int type like this
type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

//--------------------------------------------------------------------

// LEN METHOD ON STRING

type MyString string

func (s MyString) Len() int {
	return len(s)
}

//--------------------------------------------------------------------

// DELETE MULTI STRING

type MultiString []string

func (ms MultiString) Join() string {
	return strings.Join(ms, "")
}

//--------------------------------------------------------------------

// Build your own strings builder based on strings package

type MyBuilder strings.Builder

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

//--------------------------------------------------------------------

type OtherBuilder struct {
	Contents strings.Builder
}

func (ob OtherBuilder) ToUp() string {
	return strings.ToUpper(ob.Contents.String())
}

type IntBuilder struct {
	Contents []int
}

func (ib IntBuilder) Sum() int {
	total := 0

	for _, v := range ib.Contents {
		total += v
	}

	return total
}

package main

import (
	"fmt"
	"mytypes/mytypes"
)

func main() {
	var mb mytypes.MyBuilder
	fmt.Println(mb.Contents.Len())
}

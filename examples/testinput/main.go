package main

import (
	"fmt"
	"gioui.org/widget"
	"reflect"
)

func main() {
	var ed widget.Editor
	t := reflect.TypeOf(&ed)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}
}

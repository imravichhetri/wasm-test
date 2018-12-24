package main

import (
	// "fmt"
	"syscall/js"
)

func main() {
	// js.Global().Set("blackFridayFormat", js.NewCallback(format))
	// fmt.Println("Hello World!!!")
	alert := js.Global().Get("alert")
	alert.Invoke("Hello... Wasm!!!")
}

/*func format(input js.Value) {
	fmt.Println("Format Value")
}
*/

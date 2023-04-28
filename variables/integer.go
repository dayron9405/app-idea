package variables

import "fmt"

func ViewInteger() {
	var intBase int
	intOf32 := int32(10)
	intOf64 := int64(10)
	fmt.Println("intBase: ", intBase)
	fmt.Println("intOf32: ", intOf32)
	fmt.Println("intOf64: ", intOf64)
}

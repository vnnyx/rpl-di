package exception

import "fmt"

func PanicIfNeeded(err interface{}) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

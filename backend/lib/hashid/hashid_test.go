package hashid

import (
	"fmt"
	"testing"
)

func TestHashid(t *testing.T) {
	N := 5
	Len := 10

	fmt.Println("Method 1")
	for i := 0; i < N; i++ {
		fmt.Println(RandStringBytesMaskImprSrcUnsafe(Len))
	}
}

func TestUUID(t *testing.T) {
	fmt.Println(NewUUID())
}

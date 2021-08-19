package vfs_benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkTouchCopyAWS(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS("100MB_file")
	}

}




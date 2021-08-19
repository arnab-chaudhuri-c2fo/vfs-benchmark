package vfs_benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkTouchCopyAWS(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"1MB_file",
			"c2foupload-test",
			"/vfs-benchmark/folder1/",
			"stuff3.txt",
			s3Creds)
	}

}




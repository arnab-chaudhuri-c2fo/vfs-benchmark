package vfs_benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkRead_AWS_File(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		ReadRemoteFileAWS(
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"100MB_file",
			&s3Creds)
	}

}

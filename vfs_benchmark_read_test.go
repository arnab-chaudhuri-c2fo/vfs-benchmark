package vfs_benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkRead_AWS_File_100MB(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		ReadRemoteFileAWS(
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"100MB_file",
			&s3Creds)
	}

}

func BenchmarkRead_AWS_File_1GB(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		ReadRemoteFileAWS(
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"temp_1GB_file",
			&s3Creds)
	}

}

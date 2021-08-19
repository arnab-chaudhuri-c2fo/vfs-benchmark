package vfs_benchmark

import (
	"fmt"
	"testing"
)


func BenchmarkTouchCopyAWS_File250KB_Buf0(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"250KB_file",
			s3Creds, 0 * 1024)
	}

}


func BenchmarkTouchCopyAWS_File250KB_Buf256(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"250KB_file",
			s3Creds, 256 * 1024)
	}

}

func BenchmarkTouchCopyAWS_File250KB_Buf512(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"250KB_file",
			s3Creds, 512 * 1024)
	}

}

func BenchmarkTouchCopyAWS_File250KB_Buf1024(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"250KB_file",
			s3Creds, 1024 * 1024)
	}

}

//----------

func BenchmarkTouchCopyAWS_File1MB_Buf0(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"1MB_file",
			s3Creds, 0 * 1024)
	}

}


func BenchmarkTouchCopyAWS_File1MB_Buf256(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"1MB_file",
			s3Creds, 256 * 1024)
	}

}

func BenchmarkTouchCopyAWS_File1MB_Buf512(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"1MB_file",
			s3Creds, 512 * 1024)
	}

}

func BenchmarkTouchCopyAWS_File1MB_Buf1024(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"1MB_file",
			s3Creds, 1024 * 1024)
	}

}

//----------

func BenchmarkTouchCopyAWS_File10MB_Buf0(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"10MB_file",
			s3Creds, 0 * 1024)
	}

}


func BenchmarkTouchCopyAWS_File10MB_Buf256(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"10MB_file",
			s3Creds, 256 * 1024)
	}

}

func BenchmarkTouchCopyAWS_File10MB_Buf512(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"10MB_file",
			s3Creds, 512 * 1024)
	}

}

func BenchmarkTouchCopyAWS_File10MB_Buf1024(b *testing.B){

	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration - %v\n", i)
		TouchCopyAWS(
			"file:///Users/arnab.chaudhuri/Documents/C2FO/Code/vfs_benchmark/files/",
			"download",
			"c2foupload-test",
			"/vfs-benchmark/folder2/files/",
			"10MB_file",
			s3Creds, 1024 * 1024)
	}

}



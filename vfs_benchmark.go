package vfs_benchmark

import (
	"fmt"
	"github.com/c2fo/vfs/v5/backend/s3"
	"github.com/c2fo/vfs/v5/utils"
	"github.com/c2fo/vfs/v5/vfssimple"
)

//mkfile -n 10g temp_10GB_file to create dummy files

func TouchCopyAWS(fileName string){

	//Set these function consts as needed for this to work!
	const awsBucket = ""
	const remoteFilePath = ""
	const localFilePath = ""
	const accessKeyID = ""
	const secretAccessKey = ""
	const region = ""
	//

	options := 	s3.Options{
		AccessKeyID: accessKeyID,
		SecretAccessKey: secretAccessKey,
		Region: region,
	}

	s3FileSystem := s3.NewFileSystem().WithOptions(options)
	//fileA, _ := s3FileSystem.NewFile("c2foupload-test", "/vfs-benchmark/folder2/stuff2.txt")
	fileA, _ := vfssimple.NewFile( localFilePath + fileName)
	fileB, _ := s3FileSystem.NewFile(awsBucket,  remoteFilePath + fileName)



	if res, err := fileA.Exists(); err == nil{
		fmt.Printf("Starting: Filename: %v | Buffersize: %v\n", fileName, utils.BufferSize)
		fmt.Printf("File A Exists .... %v\n", res)
		cpyRes := fileA.CopyToFile(fileB)
		fmt.Printf("Copied Errors = %v\n", cpyRes)
		fmt.Println("Ending ....")
	}else{
		fmt.Printf("Error: %v\n", err)
	}


}
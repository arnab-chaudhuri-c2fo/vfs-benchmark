package vfs_benchmark

import (
	"fmt"
	"github.com/c2fo/vfs/v5/backend/s3"
	"github.com/c2fo/vfs/v5/utils"
	"github.com/c2fo/vfs/v5/vfssimple"
)

//mkfile -n 10g temp_10GB_file to create dummy files


func TouchCopyAWS(localFilePath string, localFileName string,
	remoteBucket string, remoteFilePath string,
	remoteFileName string, s3Creds S3_Creds){

	options := 	s3.Options{
		AccessKeyID: s3Creds.accessKeyID,
		SecretAccessKey: s3Creds.secretAccessKey,
		Region: s3Creds.remoteRegion,
	}

	s3FileSystem := s3.NewFileSystem().WithOptions(options)
	fileA, _ := vfssimple.NewFile( localFilePath + localFileName)
	fileB, _ := s3FileSystem.NewFile(remoteBucket,  remoteFilePath + remoteFileName)



	if res, err := fileA.Exists(); err == nil{
		fmt.Printf("Starting: Filename: %v | Buffersize: %v\n", localFileName, utils.BufferSize)
		fmt.Printf("File A Exists .... %v\n", res)
		cpyRes := fileA.CopyToFile(fileB)
		fmt.Printf("Copied Errors = %v\n", cpyRes)
		fmt.Println("Ending ....")
	}else{
		fmt.Printf("Error: %v\n", err)
	}


}
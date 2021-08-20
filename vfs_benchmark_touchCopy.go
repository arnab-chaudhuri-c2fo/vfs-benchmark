package vfs_benchmark

import (
	"fmt"
	"github.com/c2fo/vfs/v5/backend/s3"
	"github.com/c2fo/vfs/v5/vfssimple"
)

//mkfile -n 10g temp_10GB_file to create dummy files


func TouchCopyAWS(localFilePath string, localFileName string,
	remoteBucket string, remoteFilePath string,
	remoteFileName string, s3Creds *S3_Creds, fileBufferSize int){

	options := s3.Options{
		AccessKeyID: s3Creds.accessKeyID,
		SecretAccessKey: s3Creds.secretAccessKey,
		Region: s3Creds.remoteRegion,
		FileBufferSize: fileBufferSize,
	}

	s3FileSystem := s3.NewFileSystem().WithOptions(options)
	localFile, _ := vfssimple.NewFile( localFilePath + localFileName)
	remoteFile, _ := s3FileSystem.NewFile(remoteBucket, remoteFilePath + remoteFileName)



	if res, err := remoteFile.Exists(); res == true && err == nil{
		fmt.Printf("Starting: Filename: %v | Buffersize: %v\n", remoteFileName, fileBufferSize)
		fmt.Printf("\tFile A Exists .... %v\n", res)
		cpyRes := remoteFile.CopyToFile(localFile)
		fmt.Printf("\tCopied Errors = %v\n", cpyRes)
		size, _ := localFile.Size()
		fmt.Printf("\tCopied File Size = %v\n", size)
		fmt.Println("Ending ....")
	}else{
		fmt.Printf("\tFile Exists: %v | Error: %v\n", res, err)
	}


}
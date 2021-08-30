package vfs_benchmark

import (
	"fmt"
	"github.com/c2fo/vfs/v5/backend/s3"
	"io"
)

func ReadRemoteFileAWS(remoteBucket string, remoteFilePath string,
	remoteFileName string, s3Creds *S3_Creds) {

	options := s3.Options{
		AccessKeyID:     s3Creds.accessKeyID,
		SecretAccessKey: s3Creds.secretAccessKey,
		Region:          s3Creds.remoteRegion,
		FileBufferSize:  256 * 1024,
	}

	s3FileSystem := s3.NewFileSystem().WithOptions(options)
	remoteFile, _ := s3FileSystem.NewFile(remoteBucket, remoteFilePath+remoteFileName)

	data := make([]byte, 4096)
	bytesRead := 0
	for {
		data = data[:cap(data)]
		n, err := remoteFile.Read(data)
		bytesRead += n
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("BytesRead: %v\n", bytesRead)

}

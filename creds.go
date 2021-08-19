package vfs_benchmark


type S3_Creds struct{
	accessKeyID string
	secretAccessKey string
	remoteRegion string
}

//TODO Set These Values
var s3Creds = S3_Creds{
	accessKeyID: "",
	secretAccessKey: "",
	remoteRegion: "",
}
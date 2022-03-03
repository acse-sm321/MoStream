package main

import (
	"MoStream/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

// vendor is not recommended anymore
var EP string
var AK string
var SK string

func init() {
	AK = "LTAI5tFi9RV1ioteVubTUu99"       // key ID
	SK = "atO7hsqpOblypn9GBwbT5Htz7CD7D0" // key secret
	EP = config.GetOssAddr()
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("Uploading object error: %s", err)
		return false
	}

	return true
}

package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	// Load balance
	LBAddr  string `json:"lb_addr"`
	OssAddr string `json:"oss_addr"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		// must panic
		log.Printf("%v", err)
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}

func GetOssAddr() string {
	return configuration.OssAddr
}

package main

import (
	"embed"
	"encoding/json"
)

//go:embed configs/*.json
var embedFS embed.FS

// Config describes the application settings
type Config struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	TLS         bool   `json:"tls"`
	TLSCertFile string `json:"tls_cert_file"`
	TLSKeyFile  string `json:"tls_key_file"`
	Production  bool   `json:"production"`
	Database    struct {
		Host string `json:"host"`
	} `json:"database"`
}

func loadJSONConfig(fileName string) Config {
	var config Config

	bytes, err := embedFS.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		panic(err)
	}

	return config
}

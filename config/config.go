package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	System struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"system"`

	Contract struct {
		SellerAddress   string `json:"sellerAddress"`
		SellerPK        string `json:"sellerPK"`
		User2Address    string `json:"user2Address"`
		User2PK         string `json:"user2PK"`
		User3Address    string `json:"user3Address"`
		User3PK         string `json:"user3PK"`
		ContractAddress string `json:"contractAddress"`
		ChainID         int64  `json:"chainID"`
		Rpc             string `json:"rpc"`
	} `json:"contract"`
}

func CreateConfig() *Config {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &cfg
}

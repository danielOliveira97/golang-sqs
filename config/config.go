package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadAWSConfig() config.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
	)
	if err != nil {
		// handle error
		panic("Error")
	}
	return cfg
}

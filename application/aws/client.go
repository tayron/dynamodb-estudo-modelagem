package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ObterClienteDynamoDbProducao() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "sa-east-1"
		o.SharedConfigProfile = "fincapital"
		return nil
	})
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func ObterClienteDynamoDbLocal() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.SharedConfigProfile = "dynamodblocal"
		return nil
	})
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.EndpointResolver = dynamodb.EndpointResolverFromURL("http://localhost:8000")
	})
}

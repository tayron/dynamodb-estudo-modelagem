package model

import (
	"context"

	awsUtil "bitbucket/credihome/dynamodb-estudo/application/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

type Perfil struct {
	PK   string `json:"pk"`
	SK   string `json:"sk"`
	Nome string `json:"nome"`
	RG   string `json:"rg"`
	CPF  string `json:"cpf"`
}

func ObterTodasOsPerfis() []Perfil {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	listaPerfis := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName: aws.String("pedidos"),
	})

	return criarListaPerfis(listaPerfis)
}

func criarListaPerfis(listaPerfis *dynamodb.ScanPaginator) []Perfil {
	var listaPropostas []Perfil
	for listaPerfis.HasMorePages() {
		out, err := listaPerfis.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		var perfil []Perfil
		err = attributevalue.UnmarshalListOfMaps(out.Items, &perfil)
		if err != nil {
			panic(err)
		}

		listaPropostas = append(listaPropostas, perfil...)
	}
	return listaPropostas
}
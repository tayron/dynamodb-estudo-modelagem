package model

import (
	"context"

	awsUtil "github.com/tayron/dynamodb-estudo-modelagem/application/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type Telefone struct {
	Fixo    string `json:"fixo"`
	Celular string `json:"celular"`
}
type Perfil struct {
	PK        string     `json:"pk"`
	SK        string     `json:"sk"`
	Nome      string     `json:"nome"`
	RG        string     `json:"rg"`
	CPF       string     `json:"cpf"`
	Email     string     `json:"email"`
	Telefones []Telefone `json:"telefones"`
}

func ObterTodasOsPerfis() []Perfil {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	listaPerfis := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("begins_with(sk, :perfil)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":perfil": &types.AttributeValueMemberS{Value: "#PERFIL#"},
		},
	})

	return criarListaPerfis(listaPerfis)
}

func ObterListaPerfilPorSK(sk string) []Perfil {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	listaPerfis := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("sk = :sk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":sk": &types.AttributeValueMemberS{Value: sk},
		},
	})

	return criarListaPerfis(listaPerfis)
}

func ObterPerfilPorSK(sk string) Perfil {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	scanPaginator := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("sk = :sk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":sk": &types.AttributeValueMemberS{Value: sk},
		},
	})

	var perfil []Perfil
	for scanPaginator.HasMorePages() {
		out, err := scanPaginator.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		err = attributevalue.UnmarshalListOfMaps(out.Items, &perfil)
		if err != nil {
			panic(err)
		}
	}

	if len(perfil) == 0 {
		return Perfil{}
	}

	return perfil[0]
}

func criarListaPerfis(scanPaginator *dynamodb.ScanPaginator) []Perfil {
	var listaPerfis []Perfil
	for scanPaginator.HasMorePages() {
		out, err := scanPaginator.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		var perfil []Perfil
		err = attributevalue.UnmarshalListOfMaps(out.Items, &perfil)
		if err != nil {
			panic(err)
		}

		listaPerfis = append(listaPerfis, perfil...)
	}
	return listaPerfis
}

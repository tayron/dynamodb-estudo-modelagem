package model

import (
	"context"

	awsUtil "github.com/tayron/dynamodb-estudo-modelagem/application/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func ObterTodasAsFiliacoes() []Filiacao {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	resultadoConsulta := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("begins_with(sk, :perfil)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":perfil": &types.AttributeValueMemberS{Value: "#FILIACAO#"},
		},
	})

	return criarListaFiliacoes(resultadoConsulta)
}

func ObterFiliacaoPorPerfil(pk string) []Filiacao {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	resultadoConsulta := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("pk = :pk AND begins_with(sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: pk},
			":sk": &types.AttributeValueMemberS{Value: "#FILIACAO#"},
		},
	})

	return criarListaFiliacoes(resultadoConsulta)
}

func criarListaFiliacoes(resultadoConsulta *dynamodb.ScanPaginator) []Filiacao {
	var listaFiliacoes []Filiacao
	for resultadoConsulta.HasMorePages() {
		out, err := resultadoConsulta.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		var filiacao []Filiacao
		err = attributevalue.UnmarshalListOfMaps(out.Items, &filiacao)
		if err != nil {
			panic(err)
		}

		listaFiliacoes = append(listaFiliacoes, filiacao...)
	}
	return listaFiliacoes
}

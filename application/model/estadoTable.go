package model

import (
	"context"

	awsUtil "github.com/tayron/dynamodb-estudo-modelagem/application/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func ObterTodasOsEnderecos() []Endereco {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	resultadoConsulta := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("begins_with(sk, :perfil)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":perfil": &types.AttributeValueMemberS{Value: "#ENDERECO#"},
		},
	})

	return criarListaEndrecos(resultadoConsulta)
}

func ObterEnderecoPorPerfil(pk string) []Pedido {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	resultadoConsulta := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("pk = :pk AND begins_with(sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: pk},
			":sk": &types.AttributeValueMemberS{Value: "#ENDERECO#"},
		},
	})

	return criarListaPedidos(resultadoConsulta)
}

func criarListaEndrecos(resultadoConsulta *dynamodb.ScanPaginator) []Endereco {
	var listaEnderecos []Endereco
	for resultadoConsulta.HasMorePages() {
		out, err := resultadoConsulta.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		var endereco []Endereco
		err = attributevalue.UnmarshalListOfMaps(out.Items, &endereco)
		if err != nil {
			panic(err)
		}

		listaEnderecos = append(listaEnderecos, endereco...)
	}
	return listaEnderecos
}

package model

import (
	"context"

	awsUtil "github.com/tayron/dynamodb-estudo-modelagem/application/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func ObterTodasOsPedidos() []Pedido {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	resultadoConsulta := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("begins_with(sk, :perfil)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":perfil": &types.AttributeValueMemberS{Value: "#PEDIDO#"},
		},
	})

	return criarListaPedidos(resultadoConsulta)
}

func ObterPedidoPorPerfil(pk string) []Pedido {
	clienteDynamoDb := awsUtil.ObterClienteDynamoDbLocal()

	resultadoConsulta := dynamodb.NewScanPaginator(clienteDynamoDb, &dynamodb.ScanInput{
		TableName:        aws.String("pedidos"),
		FilterExpression: aws.String("pk = :pk AND begins_with(sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: pk},
			":sk": &types.AttributeValueMemberS{Value: "#PEDIDO#"},
		},
	})

	return criarListaPedidos(resultadoConsulta)
}

func criarListaPedidos(resultadoConsulta *dynamodb.ScanPaginator) []Pedido {
	var listaPedidos []Pedido
	for resultadoConsulta.HasMorePages() {
		out, err := resultadoConsulta.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		var pedido []Pedido
		err = attributevalue.UnmarshalListOfMaps(out.Items, &pedido)
		if err != nil {
			panic(err)
		}

		listaPedidos = append(listaPedidos, pedido...)
	}
	return listaPedidos
}

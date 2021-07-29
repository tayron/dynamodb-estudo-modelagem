package controller

import (
	"fmt"

	"github.com/tayron/dynamodb-estudo-modelagem/application/model"
	"github.com/tayron/dynamodb-estudo-modelagem/application/util"
)

func ObterPerfil() {
	fmt.Println("● Consultando dados no banco")
	lista := model.ObterTodasOsPerfis()

	for _, item := range lista {
		item2 := model.ObterEnderecoPorPerfil(item.PK)
		util.DebugarStruct(item2)
	}
}

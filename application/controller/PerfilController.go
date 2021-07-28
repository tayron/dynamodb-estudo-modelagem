package controller

import (
	"fmt"

	"github.com/tayron/dynamodb-estudo-modelagem/application/model"
	"github.com/tayron/dynamodb-estudo-modelagem/application/util"
)

func ObterPerfil() {
	fmt.Println()
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Obtendo todos os perfis")
	fmt.Println("---------------------------------------------------------")

	fmt.Println("● Consultando dados no banco")
	listaPerfis := model.ObterTodasOsPerfis()

	util.DebugarStruct(listaPerfis)

}

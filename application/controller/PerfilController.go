package controller

import (
	"bitbucket/credihome/dynamodb-estudo/application/model"
	"bitbucket/credihome/dynamodb-estudo/application/util"
	"fmt"
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

package model

type Perfil struct {
	PK        string `json:"pk"`
	SK        string `json:"sk"`
	Nome      string `json:"nome"`
	RG        string `json:"rg"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Telefones []struct {
		Fixo    string `json:"fixo"`
		Celular string `json:"celular"`
	} `json:"telefones"`
}

package model

type Endereco struct {
	PK          string `json:"pk"`
	SK          string `json:"sk"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"estado"`
}

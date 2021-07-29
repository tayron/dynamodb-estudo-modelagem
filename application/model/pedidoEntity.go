package model

type Pedido struct {
	PK                  string `json:"pk"`
	SK                  string `json:"sk"`
	Produto             string `json:"produto"`
	StatusEntrega       string `json:"statusEntrega"`
	Numero              string `json:"numero"`
	ValorUnitario       int    `json:"valorUnitario"`
	QuantidadeAdquirida int    `json:"quantidadeAdquirida"`
	ValorTotal          int    `json:"valorTotal"`
	StatusPedido        string `json:"statusPedido"`
}

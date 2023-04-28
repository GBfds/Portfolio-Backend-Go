package models

// types para recebimento de body
type ReqEndereco struct {
	Numero string `json:"numero"`
	Rua    string `json:"rua"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
}

// types de tabelas no DB

type Endereco struct {
	Id         string `json:"id"`
	Id_cliente string `json:"id_cliente"`
	ReqEndereco
}

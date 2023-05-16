package models

// types do gorm
type Endereco struct {
	ID         string `json:"id" gorm:"default:uuid_generate_v4()"`
	Id_cliente string `json:"id_cliente"`
	Numero     string `json:"numero"`
	Rua        string `json:"rua"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
}

// types para recebimento de body
type ReqEndereco struct {
	Numero string `json:"numero"`
	Rua    string `json:"rua"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
}

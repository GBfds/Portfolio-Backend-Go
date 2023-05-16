package models

// types do gorm
type Cliente struct {
	ID        string     `json:"id" gorm:"default:uuid_generate_v4()"`
	Nome      string     `json:"nome"`
	Email     string     `json:"email" gorm:"unique"`
	Senha     string     `json:"senha"`
	Telefone  string     `json:"telefone" gorm:"type:varchar(11)"`
	Enderecos []Endereco `gorm:"foreignKey:Id_cliente" json:"enderecos"`
}

type Admin struct {
	ID    string `json:"id" gorm:"default:uuid_generate_v4()"`
	Nome  string `json:"nome"`
	Email string `json:"email" gorm:"unique"`
	Senha string `json:"senha"`
	Cargo string `json:"carga"`
}

// types para recebimento de body
type ReqClienteSignUp struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Telefone string `json:"telefone"`
}
type ReqAdimSignUp struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Cargo string `json:"cargo"`
}

type ReqUserLogin struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

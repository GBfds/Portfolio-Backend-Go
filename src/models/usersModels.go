package models

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

// types de tabelas no DB
type Cliente struct {
	Id string `json:"id"`
	ReqClienteSignUp
}

type Admin struct {
	Id string `json:"id"`
	ReqAdimSignUp
}

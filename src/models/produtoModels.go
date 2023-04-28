package models

// types para recebimento de body
type ReqProduto struct {
	Nome    string  `json:"nome"`
	Tamanho string  `json:"tamanho"`
	Preco   float32 `json:"preco"`
}

// types de tabelas no DB
type Produto struct {
	Id string `json:"id"`
	ReqProduto
}

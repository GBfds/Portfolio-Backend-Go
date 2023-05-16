package models

// types do gorm
type Produto struct {
	ID      string  `json:"id" gorm:"default:uuid_generate_v4()"`
	Nome    string  `json:"nome"`
	Tamanho string  `json:"tamanho"`
	Preco   float32 `json:"preco"`
}

// types para recebimento de body
type ReqProduto struct {
	Nome    string  `json:"nome"`
	Tamanho string  `json:"tamanho"`
	Preco   float32 `json:"preco"`
}

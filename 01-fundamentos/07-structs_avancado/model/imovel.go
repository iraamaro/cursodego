package model

// Imovel armazena dados imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

// SetValor define valor do imóvel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

// GetValor retorna valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}

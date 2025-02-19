package model

import "errors"

// Imovel armazena dados imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

var ErrValorDeImovelInvalido = errors.New("o valor informado não é válido para o imóvel")
var ErrValorDeImovelMuitoAlto = errors.New("o valor informado é muito alto")

// SetValor define valor do imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

// GetValor retorna valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}

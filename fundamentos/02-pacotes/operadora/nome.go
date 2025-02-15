package operadora

import (
	"02-pacotes/prefixo"
	"strconv"
)

// NomeOperadora é a operadora de uso
var NomeOperadora = "XPTO Telecom"

var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) +
	" " + NomeOperadora

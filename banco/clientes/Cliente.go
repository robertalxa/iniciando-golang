package clientes

type Titular struct {
	Nome, CPF, Profissao string
	// As primeiras letras estarem maiusculas significa que a visibilidade desse campo é public
}

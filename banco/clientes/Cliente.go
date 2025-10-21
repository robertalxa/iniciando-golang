package clientes

type Titular struct {
	Nome      string
	CPF       string
	Profissao string
	// As primeiras letras estarem maiusculas significa que a visibilidade desse campo é public
}

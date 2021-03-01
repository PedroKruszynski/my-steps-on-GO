package exemplo

type Casa struct {
	Cor string
	Numero int
}

// privada
type casa struct {
	Cor string
	// propridade privada
	numero int
}

func PrintExemplo() {
	print("oiii")
}

func printExemplo() {
	print("sou uma funcao privadaaa")
}
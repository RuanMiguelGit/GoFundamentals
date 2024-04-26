package main

func main() 
{
	i := 1
	//go não tem aritimeticos de pnteiro

	var p *int = nil

	p = &i //pegando o endereço da variavel
	*p++ //desreferenciando, pegando o valor
	p++ // nao é possivel tem que fazer o procedimento acima
}
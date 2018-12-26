package main

import "fmt"
import "os"
import "net/http"

func main(){
	exibeMenu()
	entrada := lerEntrada()
	validarEntrada(entrada)
}

func validarEntrada(entrada int){
	switch entrada {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Opcao 2")
	case 0:
		fmt.Println("Adeus...")
		os.Exit(0)
	default:
		fmt.Println("Opsss!!! Nao entendi este comando")
		os.Exit(-1)
	}

}
func  lerEntrada()int{
	var comando int
	fmt.Scanf("%d", &comando)
	return comando

}
func exibeMenu(){
	fmt.Println("Digite uma opcao:")
	fmt.Println("1- Monitoramento \n2- Exibir logs \n0- Sair")
}
//fazer requisições HTTP, para fazer o monitoramento do site escolhido
func iniciarMonitoramento(){
	fmt.Println("Monitorando....")
	site := "https://www.meetup.com/pt-BR/bluetalks/events/"
	resp,_ := http.Get(site)
	fmt.Println(resp)

}